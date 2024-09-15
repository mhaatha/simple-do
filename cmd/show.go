package cmd

import (
	"github.com/mhaatha/simple-do/internal/service"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all tasks.",
	Long:  `Show all tasks from database.`,
	Run: func(cmd *cobra.Command, args []string) {
		service.ShowTasks()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
