/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// compareCmd represents the compare command
var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Comparing base benchmark %s with %v\n", args[0], args[1:])

		bindir := filepath.Join(os.Getenv("GOPATH"), "bin")
		if bindir == "bin" {
			bindir = filepath.Join(os.Getenv("HOME"), "go", "bin")
		}

		compareCmd := exec.Command(filepath.Join(bindir, "benchstat"), args...)
		compareCmd.Stdout = os.Stdout
		compareCmd.Stderr = os.Stderr
		fmt.Println(" >", compareCmd.String())
		err := compareCmd.Run()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
}
