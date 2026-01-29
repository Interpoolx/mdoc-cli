package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("mdoc-cli version %s\n", Version)
		fmt.Printf("commit: %s\n", Commit)
		fmt.Printf("build time: %s\n", BuildTime)
	},
}
