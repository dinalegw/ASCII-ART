package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	opts, err := ParseOptions(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		PrintUsage()
		os.Exit(1)
	}

	bannerPath := filepath.Join("fonts", opts.Banner+".txt")
	fontLines, err := LoadBanner(bannerPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load banner %q: %v\n", opts.Banner, err)
		os.Exit(1)
	}

	output, err := RenderASCII(opts.Text, fontLines, opts.Color)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Render error:", err)
		os.Exit(1)
	}

	if opts.OutputFile != "" {
		if err := os.WriteFile(opts.OutputFile, []byte(output), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write %q: %v\n", opts.OutputFile, err)
			os.Exit(1)
		}
		fmt.Printf("ASCII art saved to %s\n", opts.OutputFile)
		return
	}

	fmt.Print(output)
}
