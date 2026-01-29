package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint [flags] <input>",
	Short: "Lint markdown files for style and best practices",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Linting %s...\n", args[0])
		// Implementation stub
	},
}

var lintFlags struct {
	fix    bool
	rules  string
	ignore string
}

func init() {
	lintCmd.Flags().BoolVar(&lintFlags.fix, "fix", false, "Automatically fix issues")
	lintCmd.Flags().StringVar(&lintFlags.rules, "rules", "", "Rules configuration file")
	lintCmd.Flags().StringVar(&lintFlags.ignore, "ignore", "", "Patterns to ignore")

	rootCmd.AddCommand(lintCmd)
}
