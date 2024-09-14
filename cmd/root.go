package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simple-do",
	Short: "Simple-Do is a basic CLI tool to manage your to-do tasks.",
	Long:  `Simple-Do is a command-line tool for managing tasks. You can add, remove, and list tasks easily.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
