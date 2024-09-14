package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const appVersion = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Simple-Do.",
	Long:  `All software has versions. This is Simple-Do's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Simple-Do CLI v%s\n", appVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
