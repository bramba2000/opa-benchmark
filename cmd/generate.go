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

		err := generatePolicies(dir, num, names, roles, methods)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func generateString(prefix string, num int) []string {
	var result []string
	for i := 0; i < num; i++ {
		result = append(result, fmt.Sprintf("%s-%d", prefix, i))
	}
	return result
}

func generatePolicies(dir string, numPolicies int, names, roles, methods []string) error {

	err := os.MkdirAll(dir, 0755) // Create directory if it doesn't exist
	if err != nil {
		return err
	}

	for i := 0; i < numPolicies; i++ {
		name := names[i%len(names)] // Cycle through the lists
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

	return nil
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&caseName, "case", "c", "conditions", "Specify the problem name")
	generateCmd.Flags().IntVarP(&num, "num", "n", 100, "Specify the number of policies to generate")
}
