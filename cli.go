package main

import (
	"errors"
	"fmt"
	"strings"
)

// Options represents the command-line options for the ASCII art generator.
type Options struct {
	OutputFile string
	Banner     string
	Text       string
	Color      bool
}

// PrintUsage displays usage instructions for the command-line tool.
func PrintUsage() {
	fmt.Println("Usage: ascii-art-forge [options] <text>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --banner=<standard|shadow|thinkertoy>   Select the banner style (default: standard)")
	fmt.Println("  --output=<file.txt>                    Save output to a file")
	fmt.Println("  --color                                Enable colored output")
	fmt.Println("  --help                                 Show this usage guide")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ascii-art-forge Hello World")
	fmt.Println("  ascii-art-forge --banner=shadow Hello")
	fmt.Println("  ascii-art-forge --output=art.txt Hello standard")
	fmt.Println("  ascii-art-forge --banner=thinkertoy \"Hello\\nWorld\"")
	fmt.Println("  ascii-art-forge --color \"Hello World\"")
}

func ParseOptions(args []string) (Options, error) {
	if len(args) == 0 {
		return Options{}, errors.New("missing text to convert")
	}

	opts := Options{Banner: "standard"}
	textParts := make([]string, 0, len(args))

	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return Options{}, errors.New("help requested")
		}

		if strings.HasPrefix(arg, "--output=") {
			opts.OutputFile = strings.TrimPrefix(arg, "--output=")
			continue
		}

		if strings.HasPrefix(arg, "--banner=") {
			opts.Banner = strings.TrimPrefix(arg, "--banner=")
			continue
		}

		if arg == "--color" {
			opts.Color = true
			continue
		}

		textParts = append(textParts, arg)
	}

	if len(textParts) == 1 {
		opts.Text = textParts[0]
	} else if len(textParts) == 2 && isValidBanner(textParts[1]) {
		opts.Text = textParts[0]
		opts.Banner = textParts[1]
	} else {
		return Options{}, errors.New("exactly one quoted text argument is required, with optional banner after it")
	}

	if opts.OutputFile != "" && !strings.HasSuffix(opts.OutputFile, ".txt") {
		return Options{}, errors.New("output filename must end with .txt")
	}

	if !isValidBanner(opts.Banner) {
		return Options{}, fmt.Errorf("unsupported banner %q", opts.Banner)
	}

	return opts, nil
}

func isValidBanner(banner string) bool {
	switch strings.ToLower(banner) {
	case "standard", "shadow", "thinkertoy":
		return true
	default:
		return false
	}
}
