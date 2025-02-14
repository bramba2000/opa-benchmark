package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	numPolicies := flag.Int("num", 100, "Number of policies to generate")
	flag.Parse()

	names := generateStrings(100, "user")
	roles := generateStrings(100, "role")
	methods := generateStrings(100, "method")

	folder := fmt.Sprintf("./policies/conditions/%d_policies", *numPolicies)

	err := generatePolicies(folder, *numPolicies, names, roles, methods)
	if err != nil {
		fmt.Println("Error generating policies:", err)
		return
	}

	err = generateTestPolicy(folder)

	fmt.Printf("Generated %d policies in /policies/conditions\n", *numPolicies)
}

func generateStrings(count int, prefix string) []string {
	result := make([]string, count)
	for i := 0; i < count; i++ {
		result[i] = fmt.Sprintf("%s_%d", prefix, i)
	}
	return result
}

func generateTestPolicy(dir string) error {
	content := fmt.Sprintf(`package conditions_test
import data.conditions

test_allow if {
  conditions.allow with input as {"user": "user_0", "role": "role_0", "method": "method_0"}
}

test_deny if {
  not conditions.allow with input as {"user": "user_0", "role": "role_0", "method": "method_1"}
}
	`)
	filename := filepath.Join(dir, "test.rego")
	return os.WriteFile(filename, []byte(content), 0644)
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
