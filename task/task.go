package main

import (
	"cli/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "task",
		Short: "task is a CLI for managing your TODOs.",
	}
	rootCmd.AddCommand(cmd.AddCmd)
	rootCmd.AddCommand(cmd.DoCmd)
	rootCmd.AddCommand(cmd.ListCmd)
	rootCmd.Execute()
}
