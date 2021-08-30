package main

import (
	"cli/cmd"
	"cli/database"

	"github.com/spf13/cobra"
)

func main() {
	database.SetupDB()
	SetupCLI()
}

func SetupCLI() {
	var rootCmd = &cobra.Command{
		Use:   "task",
		Short: "task is a CLI for managing your TODOs.",
	}
	rootCmd.AddCommand(cmd.AddCmd)
	rootCmd.AddCommand(cmd.DoCmd)
	rootCmd.AddCommand(cmd.ListCmd)
	rootCmd.Execute()
}
