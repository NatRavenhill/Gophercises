package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DoCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do command executed")
	},
}
