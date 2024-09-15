package cmd

import (
	"fmt"
	"strconv"

	"github.com/mhaatha/simple-do/internal/service"
	"github.com/spf13/cobra"
)

var id uint16
var descriptionUpdate string
var isDone string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task.",
	Long:  `Update description or status of a task.`,
	Run: func(cmd *cobra.Command, args []string) {
		if isDone != "true" && isDone != "false" {
			fmt.Println("Error: -s or --status only takes 'true' or 'false'")
			cmd.Usage()
			return
		}
		if id == 0 {
			fmt.Println("Error: id is required. Use -i or --id to specify the task id.")
			cmd.Usage()
			return
		}

		bool, err := strconv.ParseBool(isDone)
		if err != nil {
			fmt.Println("Error:", err)
		}

		service.UpdateTask(id, descriptionUpdate, bool)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().Uint16VarP(&id, "id", "i", 0, "Target task id to update. (required)")
	updateCmd.Flags().StringVarP(&descriptionUpdate, "update-description", "u", "", "Update to-do list's description. Should use quotation marks if the data contains spaces.")
	updateCmd.Flags().StringVarP(&isDone, "status", "s", "", "Update to-do list's status. Only takes 'true' or 'false'.")
}
