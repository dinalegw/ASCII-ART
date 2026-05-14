package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
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

// applyColor applies a random color to each character in the ASCII art.
func applyColor(asciiArt string, enableColor bool) string {
	if !enableColor {
		return asciiArt
	}

	lines := strings.Split(asciiArt, "\n")
	var coloredLines []string

	for _, line := range lines {
		var coloredLine strings.Builder
		for _, char := range line {
			if char == ' ' {
				coloredLine.WriteRune(char)
			} else {
				color := colorCodes[rand.Intn(len(colorCodes))]
				coloredLine.WriteString(color)
				coloredLine.WriteRune(char)
				coloredLine.WriteString(resetColor)
			}
		}
		coloredLines = append(coloredLines, coloredLine.String())
	}

	return strings.Join(coloredLines, "\n")
}

// RenderASCII converts input text into ASCII art using the provided banner.
func RenderASCII(input string, bannerLines []string, enableColor bool) (string, error) {
	text := strings.ReplaceAll(input, "\\n", "\n")
	lines := strings.Split(text, "\n")

	if len(lines) == 0 {
		return "", nil
	}

	const lineHeight = 8
	const charBlockSize = 9

	var output strings.Builder

	for lineIndex, line := range lines {
		if line == "" {
			if lineIndex > 0 {
				output.WriteString("\n")
			}
			output.WriteString("\n")
			continue
		}

		for row := 0; row < lineHeight; row++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					continue
				}

				index := (int(char)-32)*charBlockSize + 1
				if index+row >= len(bannerLines) {
					return "", fmt.Errorf("banner file is malformed or incomplete")
				}

				output.WriteString(bannerLines[index+row])
			}
			output.WriteString("\n")
		}
	}

	return applyColor(output.String(), enableColor), nil
}
