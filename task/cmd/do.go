package cmd

import (
	"cli/database"
	"log"

	"github.com/spf13/cobra"
)

var DoCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("No task given!")
		}

		database.CompleteTask(args[0])
	},
}
