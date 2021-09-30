package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migration called")
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}
