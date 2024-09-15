package cmd

import (
	"fmt"

	"github.com/mhaatha/simple-do/internal/service"
	"github.com/spf13/cobra"
)

var idDelete uint16

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task.",
	Long:  `Delete a task from database.`,
	Run: func(cmd *cobra.Command, args []string) {
		if idDelete == 0 {
			fmt.Println("Error: id is required. Use -i or --id to specify the task id.")
			cmd.Usage()
			return
		}

		service.DeleteTask(idDelete)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Uint16VarP(&idDelete, "id", "i", 0, "Target task id to delete. (required)")
}
