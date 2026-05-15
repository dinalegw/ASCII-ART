package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Options represents the command-line options for the ASCII art generator.
type Options struct {
	OutputFile  string
	Banner      string
	Text        string
	Color       string
	Justify     string
	Interactive bool
}

// PrintUsage displays usage instructions for the command-line tool.
func PrintUsage() {
	fmt.Println("Usage: ascii-art-forge [options] <text>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --banner=<standard|shadow|thinkertoy>   Select the banner style (default: standard)")
	fmt.Println("  --output=<file.txt>                    Save output to a file")
	fmt.Println("  --color=<red|green|yellow|blue|magenta|cyan|random>  Color mode (default: random)")
	fmt.Println("  --justify=<left|center|right>           Text alignment (default: left)")
	fmt.Println("  --help, -h                             Show this usage guide")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ascii-art-forge Hello World")
	fmt.Println("  ascii-art-forge --banner=shadow Hello")
	fmt.Println("  ascii-art-forge --output=art.txt Hello standard")
	fmt.Println("  ascii-art-forge --banner=thinkertoy \"Hello\\nWorld\"")
	fmt.Println("  ascii-art-forge --color=blue \"Hello World\"")
	fmt.Println("  ascii-art-forge --justify=center \"Hello World\"")
	fmt.Println("  ascii-art-forge                         (interactive mode)")
}

// colorNames maps color names to ANSI codes
var colorNames = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
}

func ParseOptions(args []string) (Options, error) {
	opts := Options{Banner: "standard", Color: "random", Justify: "left"}
	textParts := make([]string, 0, len(args))

	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return Options{}, fmt.Errorf("help requested")
		}

		if strings.HasPrefix(arg, "--output=") {
			opts.OutputFile = strings.TrimPrefix(arg, "--output=")
			continue
		}

		if strings.HasPrefix(arg, "--banner=") {
			opts.Banner = strings.TrimPrefix(arg, "--banner=")
			continue
		}

		if strings.HasPrefix(arg, "--color=") {
			opts.Color = strings.TrimPrefix(arg, "--color=")
			continue
		}

		if arg == "--color" {
			opts.Color = "random"
			continue
		}

		if strings.HasPrefix(arg, "--justify=") {
			opts.Justify = strings.TrimPrefix(arg, "--justify=")
			continue
		}

		textParts = append(textParts, arg)
	}

	if len(textParts) == 0 {
		return opts, nil
	}

	if len(textParts) == 1 {
		opts.Text = textParts[0]
	} else if len(textParts) == 2 && isValidBanner(textParts[1]) {
		opts.Text = textParts[0]
		opts.Banner = textParts[1]
	} else {
		return Options{}, fmt.Errorf("usage: provide text, optionally followed by banner name (standard|shadow|thinkertoy)")
	}

	if opts.OutputFile != "" && !strings.HasSuffix(opts.OutputFile, ".txt") {
		return Options{}, fmt.Errorf("output filename must end with .txt")
	}

	if !isValidBanner(opts.Banner) {
		return Options{}, fmt.Errorf("unsupported banner %q", opts.Banner)
	}

	color := strings.ToLower(opts.Color)
	if color != "" && color != "random" && !isValidColor(color) {
		return Options{}, fmt.Errorf("unsupported color %q (use red, green, yellow, blue, magenta, cyan, or random)", opts.Color)
	}

	justify := strings.ToLower(opts.Justify)
	if !isValidJustify(justify) {
		return Options{}, fmt.Errorf("unsupported justify %q (use left, center, or right)", opts.Justify)
	}
	opts.Justify = justify

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

func isValidColor(color string) bool {
	_, ok := colorNames[color]
	return ok
}

func isValidJustify(justify string) bool {
	switch strings.ToLower(justify) {
	case "left", "center", "right":
		return true
	default:
		return false
	}
}

func runInteractiveMode() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		fmt.Fprintln(os.Stderr, "No text provided")
		os.Exit(1)
	}

	fmt.Print("Banner (standard/shadow/thinkertoy) [standard]: ")
	banner, _ := reader.ReadString('\n')
	banner = strings.TrimSpace(banner)
	if banner == "" {
		banner = "standard"
	}
	if !isValidBanner(banner) {
		fmt.Fprintf(os.Stderr, "Invalid banner: %s\n", banner)
		os.Exit(1)
	}

	fmt.Print("Color (red/green/yellow/blue/magenta/cyan/random) [random]: ")
	color, _ := reader.ReadString('\n')
	color = strings.ToLower(strings.TrimSpace(color))
	if color == "" {
		color = "random"
	}
	if !isValidColor(color) && color != "random" {
		fmt.Fprintf(os.Stderr, "Invalid color: %s\n", color)
		os.Exit(1)
	}

	fmt.Print("Justify (left/center/right) [left]: ")
	justify, _ := reader.ReadString('\n')
	justify = strings.ToLower(strings.TrimSpace(justify))
	if justify == "" {
		justify = "left"
	}
	if !isValidJustify(justify) {
		fmt.Fprintf(os.Stderr, "Invalid justify: %s\n", justify)
		os.Exit(1)
	}

	opts := Options{Text: text, Banner: banner, Color: color, Justify: justify}
	run(opts)
}

func run(opts Options) {
	bannerPath := filepath.Join("fonts", opts.Banner+".txt")
	fontLines, err := LoadBanner(bannerPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load banner %q: %v\n", opts.Banner, err)
		os.Exit(1)
	}

	output, err := RenderASCII(opts.Text, fontLines, opts.Color, opts.Justify)
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

func main() {
	opts, err := ParseOptions(os.Args[1:])
	if err != nil {
		if strings.Contains(err.Error(), "help requested") {
			PrintUsage()
			os.Exit(0)
		}
		fmt.Fprintln(os.Stderr, err)
		if strings.Contains(err.Error(), "missing text") || strings.Contains(err.Error(), "usage:") {
			runInteractiveMode()
			return
		}
		PrintUsage()
		os.Exit(1)
	}

	if opts.Text == "" {
		runInteractiveMode()
		return
	}

	run(opts)
}
