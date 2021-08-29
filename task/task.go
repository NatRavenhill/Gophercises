package main

import (
	"cli/cmd"
	"cli/database"

	"github.com/spf13/cobra"
)

func main() {
	//SetupCLI()

	//database testing
	database.SetupDB()
	database.AddEntry("placeholder text")
	database.AddEntry("second placeholder text")
	database.ShowEntries()
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
