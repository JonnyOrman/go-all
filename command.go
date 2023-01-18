package main

import "github.com/spf13/cobra"

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
