/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test <case>",
	Short: "A brief description of your command",
	Long: `Run a benchmark test cases and store results in /results folder. 

You can specify one of the available benchmark case setting <case> argument. 
Also, you can specify the number of policies to generate for each case using --num flag.`,
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgs: availableCases,
	Run: func(cmd *cobra.Command, args []string) {
		// Load flags and arguments
		caseName := args[0]
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			panic(err)
		}
		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			panic(err)
		}

		// Setup the test command
		testCmd := exec.Command("opa", "test", fmt.Sprintf("policies/%s/%d", caseName, num), "-v", "--bench",
			"--format", "gobench", "--count", strconv.Itoa(count))
		fmt.Println(" >", testCmd.String())

		// Create output directory and file
		err = os.MkdirAll(fmt.Sprintf("results/%s", caseName), 0755)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(fmt.Sprintf("results/%s/%d.txt", caseName, num))
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Run the test command
		testCmd.Stdout = file
		err = testCmd.Run()
		if err != nil {
			panic(err)
		}

		fmt.Println("Results saved in results folder")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().IntP("num", "n", 10, "Number of policies to generate")
	testCmd.Flags().IntP("count", "c", 1, "Number of times to run the test")
}
