package cmd

import (
	"github.com/interpoolx/mdoc-cli/internal/converter"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert [flags] <input>",
	Short: "Convert markdown to other formats",
	Long: `Convert markdown files to HTML, PDF, DOCX, or JSON formats.
    
Supports batch conversion with concurrent processing for improved performance.`,
	Args: cobra.ExactArgs(1),
	RunE: runConvert,
}

var convertFlags struct {
	to        string
	output    string
	template  string
	recursive bool
	workers   int
	css       string
	theme     string
}

func init() {
	convertCmd.Flags().StringVarP(&convertFlags.to, "to", "t", "html",
		"target format (html, pdf, docx, json)")
	convertCmd.Flags().StringVarP(&convertFlags.output, "output", "o", "",
		"output path")
	convertCmd.Flags().StringVar(&convertFlags.template, "template", "",
		"custom template path")
	convertCmd.Flags().BoolVarP(&convertFlags.recursive, "recursive", "r", false,
		"process directories recursively")
	convertCmd.Flags().IntVarP(&convertFlags.workers, "workers", "w", 4,
		"concurrent workers (1-32)")
	convertCmd.Flags().StringVar(&convertFlags.css, "css", "",
		"custom CSS file (HTML/PDF)")
	convertCmd.Flags().StringVar(&convertFlags.theme, "theme", "default",
		"output theme")

	rootCmd.AddCommand(convertCmd)
}

func runConvert(cmd *cobra.Command, args []string) error {
	input := args[0]

	conv := converter.New(converter.Options{
		Format:    convertFlags.to,
		Output:    convertFlags.output,
		Template:  convertFlags.template,
		CSS:       convertFlags.css,
		Theme:     convertFlags.theme,
		Workers:   convertFlags.workers,
		Recursive: convertFlags.recursive,
	})

	return conv.Convert(input)
}
