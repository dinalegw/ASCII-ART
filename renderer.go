package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const (
	lineHeight    = 8
	charBlockSize = 9
)

// LoadBanner reads a font file and returns its lines.
func LoadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

// colorCodes is a list of ANSI color codes for coloring ASCII art.
var colorCodes = []string{
	"\033[31m", // Red
	"\033[32m", // Green
	"\033[33m", // Yellow
	"\033[34m", // Blue
	"\033[35m", // Magenta
	"\033[36m", // Cyan
}

// resetColor resets the color back to default.
const resetColor = "\033[0m"

// applyColor applies color to each character in the ASCII art.
// If color is "random", each character gets a random color.
// Otherwise, uses the specified color.
func applyColor(asciiArt string, color string) string {
	lines := strings.Split(asciiArt, "\n")
	var coloredLines []string

	for _, line := range lines {
		var coloredLine strings.Builder
		for _, char := range line {
			if char == ' ' {
				coloredLine.WriteRune(char)
				continue
			}

			var colorCode string
			if color == "random" {
				colorCode = colorCodes[rand.Intn(len(colorCodes))]
			} else if c, ok := colorNames[color]; ok {
				colorCode = c
			} else {
				colorCode = colorCodes[rand.Intn(len(colorCodes))]
			}

			coloredLine.WriteString(colorCode)
			coloredLine.WriteRune(char)
			coloredLine.WriteString(resetColor)
		}
		coloredLines = append(coloredLines, coloredLine.String())
	}

	return strings.Join(coloredLines, "\n")
}

// getMaxWidth returns the maximum line width in the ASCII art.
func getMaxWidth(asciiArt string) int {
	lines := strings.Split(asciiArt, "\n")
	maxWidth := 0
	for _, line := range lines {
		// Count non-color characters
		cleanLine := stripAnsi(line)
		if len(cleanLine) > maxWidth {
			maxWidth = len(cleanLine)
		}
	}
	return maxWidth
}

// stripAnsi removes ANSI escape codes from a string.
func stripAnsi(s string) string {
	var result strings.Builder
	inEscape := false
	for _, r := range s {
		if r == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape {
			if r == 'm' {
				inEscape = false
			}
			continue
		}
		result.WriteRune(r)
	}
	return result.String()
}

// justifyLine applies alignment to a single line.
func justifyLine(line, justify string, width int) string {
	cleanLine := stripAnsi(line)
	if justify == "right" {
		return strings.Repeat(" ", width-len(cleanLine)) + line
	} else if justify == "center" {
		padding := (width - len(cleanLine)) / 2
		return strings.Repeat(" ", padding) + line
	}
	return line
}

// justifyASCII applies alignment to the entire ASCII art.
func justifyASCII(asciiArt string, justify string) string {
	lines := strings.Split(asciiArt, "\n")
	if justify == "left" {
		return asciiArt
	}

	maxWidth := getMaxWidth(asciiArt)
	var result strings.Builder
	for i, line := range lines {
		if i > 0 {
			result.WriteString("\n")
		}
		result.WriteString(justifyLine(line, justify, maxWidth))
	}
	return result.String()
}

// RenderASCII converts input text into ASCII art using the provided banner.
func RenderASCII(input string, bannerLines []string, color, justify string) (string, error) {
	if len(bannerLines) == 0 {
		return "", fmt.Errorf("empty banner file")
	}

	text := strings.ReplaceAll(input, "\\n", "\n")
	lines := strings.Split(text, "\n")

	if len(lines) == 0 {
		return "", nil
	}

	var output strings.Builder
	prevLineWasEmpty := false

	for lineIndex, line := range lines {
		if line == "" {
			if prevLineWasEmpty && lineIndex > 0 {
				output.WriteString("\n")
			}
			prevLineWasEmpty = true
			continue
		}

		prevLineWasEmpty = false

		for row := 0; row < lineHeight; row++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					continue
				}

				index := (int(char)-32)*charBlockSize + 1
				if index+row >= len(bannerLines) {
					return "", fmt.Errorf("banner file is malformed or incomplete: character %q at position %d", char, index+row)
				}

				output.WriteString(bannerLines[index+row])
			}
			output.WriteString("\n")
		}
	}

	result := applyColor(output.String(), color)
	return justifyASCII(result, justify), nil
}
