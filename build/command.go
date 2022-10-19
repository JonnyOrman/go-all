package build

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO - build")
	},
}

func AddTo(parentCommand *cobra.Command) {
	parentCommand.AddCommand(buildCmd)
}
