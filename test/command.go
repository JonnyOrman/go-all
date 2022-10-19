package test

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO - test")
	},
}

func AddTo(parentCommand *cobra.Command) {
	parentCommand.AddCommand(testCmd)
}
