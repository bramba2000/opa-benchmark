/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

var caseName string
var num int

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a test case",
	Long: `Generate a test case for the given problem. 

You can specify one of the available problems name using --case flag. 
Also, you can specify the number of policies to generate for each test cases using --num flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir := filepath.Join("policies", caseName, strconv.Itoa(num))
		names := generateString("user", num)
		roles := generateString("role", num)
		methods := generateString("method", num)

		var err error
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
	r := rand.Int() % numPolicies
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
	generateCmd.Flags().StringVarP(&caseName, "case", "c", "conditions", "Specify the problem name")
	generateCmd.Flags().IntVarP(&num, "num", "n", 100, "Specify the number of policies to generate")
}
