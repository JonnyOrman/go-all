package main

import (
	"fmt"
	"os"

	"github.com/jonnyorman/go-all/build"
	"github.com/jonnyorman/go-all/mod"
	"github.com/jonnyorman/go-all/test"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:     "go-all",
	Version: version,
	Short:   "go-all - perform command line actions on all go modules in subdirectories of the target directory",
	Long: `perform command line actions on all go modules in subdirectories of the target directory,
   
indluding:
- building
- tidying
- testing`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	build.AddTo(rootCmd)

	mod.AddTo(rootCmd)

	test.AddTo(rootCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing go-all '%s'", err)
		os.Exit(1)
	}
}
