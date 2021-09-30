package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpc called")
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
