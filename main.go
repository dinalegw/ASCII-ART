package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		fmt.Fprintln(os.Stderr, "Error: No text provided")
		os.Exit(1)
	}

	fmt.Print("Banner (standard/shadow/thinkertoy) [standard]: ")
	banner, _ := reader.ReadString('\n')
	banner = strings.TrimSpace(banner)
	if banner == "" {
		banner = "standard"
	}

	if !isValidBanner(banner) {
		fmt.Fprintf(os.Stderr, "Error: Invalid banner %q\n", banner)
		os.Exit(1)
	}

	fmt.Print("Color (red/green/yellow/blue/magenta/cyan/random) [random]: ")
	color, _ := reader.ReadString('\n')
	color = strings.ToLower(strings.TrimSpace(color))
	if color == "" {
		color = "random"
	}

	if !isValidColor(color) {
		fmt.Fprintf(os.Stderr, "Error: Invalid color %q\n", color)
		os.Exit(1)
	}

	fmt.Print("Justify (left/center/right/justify) [left]: ")
	justify, _ := reader.ReadString('\n')
	justify = strings.ToLower(strings.TrimSpace(justify))
	if justify == "" {
		justify = "left"
	}

	if !isValidJustify(justify) {
		fmt.Fprintf(os.Stderr, "Error: Invalid justify %q\n", justify)
		os.Exit(1)
	}

	bannerPath := filepath.Join("fonts", banner+".txt")
	fontLines, err := LoadBanner(bannerPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to load banner: %v\n", err)
		os.Exit(1)
	}

	output, err := RenderASCII(text, fontLines, color, justify)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Render failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(output)
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
	switch color {
	case "red", "green", "yellow", "blue", "magenta", "cyan", "random":
		return true
	default:
		return false
	}
}

func isValidJustify(justify string) bool {
	switch justify {
	case "left", "center", "right", "justify": // ✅ FIX HERE
		return true
	default:
		return false
	}
}