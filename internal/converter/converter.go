package converter

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
)

type Options struct {
	Format    string
	Output    string
	Template  string
	CSS       string
	Theme     string
	Workers   int
	Recursive bool
}

type Converter struct {
	opts Options
}

func New(opts Options) *Converter {
	return &Converter{opts: opts}
}

func (c *Converter) Convert(input string) error {
	info, err := os.Stat(input)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return c.convertDir(input)
	}

	return c.convertFile(input)
}

func (c *Converter) convertFile(input string) error {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		return err
	}

	var output []byte
	switch c.opts.Format {
	case "html":
		output = markdown.ToHTML(data, nil, nil)
	case "json":
		// Mock JSON AST for now
		output = []byte(fmt.Sprintf(`{"file": "%s", "type": "markdown"}`, input))
	default:
		return fmt.Errorf("unsupported format: %s", c.opts.Format)
	}

	if c.opts.Output == "" {
		fmt.Print(string(output))
		return nil
	}

	return ioutil.WriteFile(c.opts.Output, output, 0644)
}

func (c *Converter) convertDir(input string) error {
	return filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			// In a real implementation, this would use workers
			return c.convertFile(path)
		}
		if info.IsDir() && path != input && !c.opts.Recursive {
			return filepath.SkipDir
		}
		return nil
	})
}
