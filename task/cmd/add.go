package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add command executed")
	},
}
