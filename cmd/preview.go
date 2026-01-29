package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var previewCmd = &cobra.Command{
	Use:   "preview [flags] <input>",
	Short: "Preview markdown as HTML in browser",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Previewing %s...\n", args[0])
		// Implementation stub
	},
}

var previewFlags struct {
	port  int
	watch bool
	theme string
}

func init() {
	previewCmd.Flags().IntVar(&previewFlags.port, "port", 8080, "Server port")
	previewCmd.Flags().BoolVar(&previewFlags.watch, "watch", false, "Watch for changes and reload")
	previewCmd.Flags().StringVar(&previewFlags.theme, "theme", "default", "Preview theme")

	rootCmd.AddCommand(previewCmd)
}
