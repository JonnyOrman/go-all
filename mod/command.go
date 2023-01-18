package mod

import (
	"fmt"

	"github.com/jonnyorman/go-all/mod/tidy"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:   "mod",
	Short: "mod commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO - mod")
	},
}

func AddTo(parentCommand *cobra.Command) {
	tidy.AddTo(command)
	parentCommand.AddCommand(command)
}
