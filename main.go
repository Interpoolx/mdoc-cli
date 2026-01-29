package main

import (
	"os"

	"github.com/interpoolx/mdoc-cli/cmd"
)

var (
	Version   = "dev"
	Commit    = "unknown"
	BuildTime = "unknown"
)

func main() {
	cmd.Version = Version
	cmd.Commit = Commit
	cmd.BuildTime = BuildTime

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
