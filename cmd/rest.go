package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rest called")
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
