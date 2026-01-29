package ui

import (
	"fmt"

	"github.com/fatih/color"
)

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

func PrintBanner() {
	banner := `
  __  __   _____                 
 |  \/  | |  __ \                
 | \  / | | |  | |  ___    ___   
 | |\/| | | |  | | / _ \  / __|  
 | |  | | | |__| || (_) || (__   
 |_|  |_| |_____/  \___/  \___|  
                                 
                                 `
	color.HiYellow(banner)
	color.HiBlack("  Markdown Document CLI - High Performance Converter")
	color.HiBlack("  Convert your markdown files to beautiful HTML pages.\n\n")
}

func ShowAppHeader(version, author string) {
	PrintBanner()
	PrintVersion(version, author)
}

func PrintVersion(version, author string) {
	Info.Print("MDoc CLI ")
	Success.Println(version)
	fmt.Println()

	fmt.Printf("Author:    %s\n", author)
	fmt.Printf("Portfolio: %s\n", "web4strategy.com")
	fmt.Printf("Twitter/X: %s\n", "@web4strategy")

	fmt.Println()
	Info.Println("NPM Package Info")
	fmt.Printf("  Package: %-10s\n", "mdoc-cli")
	fmt.Printf("  Latest:  %-10s\n", "npm install -g mdoc-cli@latest")
	fmt.Println()
}

func PrintConversionInfo(input, output, format string) {
	Info.Println("Conversion Summary")
	fmt.Printf("  Input:  %s\n", input)
	fmt.Printf("  Output: %s\n", output)
	fmt.Printf("  Format: %s\n", format)
	fmt.Println()
}
