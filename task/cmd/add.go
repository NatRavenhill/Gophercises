package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"

	"cli/database"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("No task provided!")
		}

		database.AddTask(strings.Join(args[0:], " "))
	},
}
