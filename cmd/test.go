package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Run:   runTestCmd,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func runTestCmd(cmd *cobra.Command, args []string) {
	// TODO: validate args length

	intValue, err := strconv.Atoi(args[0])
	if nil != err {
		// TODO: handle error
	}

	for i := 1; i <= intValue; i++ {
		str := ""

		if i % 3 == 0 {
			str = "Chia het cho 3"
		}

		if i % 5 == 0 {
			str = "chia het cho 5"
		}

		if "" == str {
			str = fmt.Sprintf("%d", i)
		}

		fmt.Println(str)
	}
}
