package ui

import "github.com/fatih/color"

var (
	Success = color.New(color.FgGreen)
	Error   = color.New(color.FgRed)
	Warning = color.New(color.FgYellow)
	Info    = color.New(color.FgCyan)
)

func PrintSuccess(msg string) {
	Success.Println("✓ " + msg)
}

func PrintError(msg string) {
	Error.Println("✗ " + msg)
}

func PrintInfo(msg string) {
	Info.Println("→ " + msg)
}
