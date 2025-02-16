/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate <case>",
	Short: "Generate a test case",
	Long: `Generate a test case for the given problem. 

You can specify one of the available benchmark case setting <case> argument. 
Also, you can specify the number of policies to generate for each case using --num flag.`,
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgs: availableCases,
	Run: func(cmd *cobra.Command, args []string) {
		caseName := args[0]
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			fmt.Println(err)
			return
		}

		dir := filepath.Join("policies", caseName, strconv.Itoa(num))
		names := generateString("user", num)
		roles := generateString("role", num)
		methods := generateString("method", num)

		switch caseName {
		case "conditions":
			err = generateConditionsCase(dir, num, names, roles, methods)
		default:
			err = fmt.Errorf("unknown case name: %s", caseName)
		}

		if err != nil {
			fmt.Println(err)
		}
	},
}

func generateString(prefix string, num int) []string {
	var result []string
	for i := 0; i < num; i++ {
		result = append(result, fmt.Sprintf("%s_%d", prefix, i))
	}
	return result
}

func generateConditionsCase(dir string, numPolicies int, names, roles, methods []string) error {

	err := os.MkdirAll(dir, 0755) // Create directory if it doesn't exist
	if err != nil {
		return err
	}

	// Generate the policies
	for i := 0; i < numPolicies; i++ {
		name := names[i%len(names)]
		role := roles[i%len(roles)]
		method := methods[i%len(methods)]

		filename := filepath.Join(dir, fmt.Sprintf("policy_%d.rego", i))

		policy := fmt.Sprintf(`package conditions

allow if {
  input.user == "%s"
  input.role == "%s"
  input.method == "%s"
}`, name, role, method)

		err := os.WriteFile(filename, []byte(policy), 0644)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Generated %d policies in %s\n", numPolicies, dir)

	// Genrate test case
	filename := filepath.Join(dir, "test.rego")
	r := numPolicies / 2
	test := fmt.Sprintf(`package conditions_test
import data.conditions

test_allow if {
  conditions.allow with input as {"user": "user_%d", "role": "role_%d", "method": "method_%d"}
}

test_deny if {
  not conditions.allow with input as {"user": "user_%d", "role": "role_%d", "method": "method_%d"}
}
	`, r, r, r, r, r, r+1)

	err = os.WriteFile(filename, []byte(test), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Generated test case in %s with random test value %d\n", filename, r)

	return nil
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("num", "n", 10, "Number of policies to generate")
}
