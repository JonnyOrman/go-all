package test

import (
	"os"
	"path/filepath"

	"github.com/jonnyorman/go-all/command"
	"github.com/spf13/cobra"
)

var goAllCommand = command.GenerateCommand("Testing", "test")

var cobraCommand = &cobra.Command{
	Use:   "test",
	Short: "Test commands",
	Run: func(cmd *cobra.Command, args []string) {
		ex, _ := os.Executable()
		directory := filepath.Dir(ex)

		goAllCommand.ExecuteAllIn(directory)
	},
}

func AddTo(parentCommand *cobra.Command) {
	parentCommand.AddCommand(cobraCommand)
}
