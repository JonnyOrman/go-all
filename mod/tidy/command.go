package tidy

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "tidy commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO - tidy")
	},
}

func AddTo(parentCommand *cobra.Command) {
	parentCommand.AddCommand(tidyCmd)
}
