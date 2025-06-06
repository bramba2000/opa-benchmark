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
		case "array":
			err = generateArrayCase(dir, num, names, roles, methods)
		case "early_exit":
			err = generateEarlyExit(dir, num, names, roles, methods)
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
	r := rand.Intn(numPolicies)
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

func generateArrayCase(dir string, numPolicies int, names, roles, methods []string) error {
	err := os.MkdirAll(dir, 0755) // Create directory if it doesn't exist
	if err != nil {
		return err
	}

	// Generate the policies
	var content string = "package arrays\n\nroles := [\n"
	for i := 0; i < numPolicies; i++ {
		content += fmt.Sprintf(`%c{"%s":"%s", "%s":"%s", "%s":"%s"}%s`, '\t', "user", names[i%len(names)], "role", roles[i%len(roles)], "method", methods[i%len(methods)], ",\n")
	}
	content += "]\n"
	filename := filepath.Join(dir, "policy.rego")
	err = os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Generated %d policies in %s\n", numPolicies, dir)

	// Genrate test case
	filename = filepath.Join(dir, "test.rego")
	r := numPolicies / 2
	test := fmt.Sprintf(`package arrays_test
import data.arrays

test_allow if {
	some a in data.arrays[_]
	a.role == "%s"
}
`, "role_"+strconv.Itoa(r))
	err = os.WriteFile(filename, []byte(test), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Generated test case in %s with random test value %d\n", filename, r)
	return nil
}

func generateEarlyExit(dir string, numPolicies int, names, roles, methods []string) error {
	err := os.MkdirAll(dir, 0755) // Create directory if it doesn't exist
	if err != nil {
		return err
	}

	// Generate the policies
	content := `package early_exit` + "\n\n"
	for i := 0; i < numPolicies; i++ {
		name := names[i%len(names)]
		role := roles[i%len(roles)]
		method := methods[i%len(methods)]

		content += "allow if {"
		r := rand.Intn(2)
		if r == 0 {
			content += fmt.Sprintf(`%cinput.user == "%s"%c`, '\t', name, '\n')
		}
		r = rand.Intn(2)
		if r == 0 {
			content += fmt.Sprintf(`%cinput.role == "%s"%c`, '\t', role, '\n')
		}
		r = rand.Intn(2)
		if r == 0 {
			content += fmt.Sprintf(`%cinput.method == "%s"%c`, '\t', method, '\n')
		}
		r = rand.Intn(2)
		if r == 0 {
			content += fmt.Sprintf(`%cinput.userrole == "%s"%c`, '\t', name+role, '\n')
		}
		r = rand.Intn(2)
		if r == 0 {
			content += fmt.Sprintf(`%cinput.usermethod == "%s"%c`, '\t', name+method, '\n')
		}
		r = rand.Intn(2)
		if r == 0 {
			content += fmt.Sprintf(`%cinput.rolemethod == "%s"%c`, '\t', role+method, '\n')
		}
		content += "}\n\n"
	}

	filename := filepath.Join(dir, "policy.rego")
	err = os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Generated %d policies in %s\n", numPolicies, dir)

	// Generate test case
	filename = filepath.Join(dir, "test.rego")
	r := rand.Intn(numPolicies)
	test := fmt.Sprintf(`package early_exit_test
import data.early_exit

test_allow if {
  early_exit.allow with input as {"user": "user_%d", "role": "role_%d", "method": "method_%d", "userrole": "user_%drole_%d", "usermethod": "user_%dmethod_%d", "rolemethod": "role_%dmethod_%d"}
}
	`, r, r, r, r, r, r, r, r, r)

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
