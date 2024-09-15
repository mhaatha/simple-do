package cmd

import (
	"fmt"

	"github.com/mhaatha/simple-do/internal/service"
	"github.com/spf13/cobra"
)

var description string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task.",
	Long:  `Insert a new task to database.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(description) == 0 {
			fmt.Println("Error: description is required. Use -d or --description to specify the task description.")
			cmd.Usage()
			return
		}

		service.AddTask(description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&description, "description", "d", "", "To-do list's description. Should use quotation marks if the data contains spaces. (required)")
}
