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

var colorMap = map[string]string{
	"red":     "\x1b[31m",
	"green":   "\x1b[32m",
	"yellow":  "\x1b[33m",
	"blue":    "\x1b[34m",
	"magenta": "\x1b[35m",
	"cyan":    "\x1b[36m",
}

var randomColors = []string{
	"\x1b[31m",
	"\x1b[32m",
	"\x1b[33m",
	"\x1b[34m",
	"\x1b[35m",
	"\x1b[36m",
}

const resetCode = "\x1b[0m"

func LoadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func applyColor(asciiArt string, color string) string {
	lines := strings.Split(asciiArt, "\n")
	var result strings.Builder

	for _, line := range lines {
		for _, ch := range line {
			if ch == ' ' {
				result.WriteRune(ch)
				continue
			}

			var code string
			if color == "random" {
				code = randomColors[rand.Intn(len(randomColors))]
			} else {
				code = colorMap[color]
			}

			result.WriteString(code)
			result.WriteRune(ch)
			result.WriteString(resetCode)
		}
		result.WriteString("\n")
	}

	return result.String()
}

func justifyASCII(asciiArt string, justify string) string {
	lines := strings.Split(asciiArt, "\n")
	if justify == "left" || len(lines) == 0 {
		return asciiArt
	}

	maxWidth := 0
	for _, line := range lines {
		clean := stripAnsi(line)
		if len(clean) > maxWidth {
			maxWidth = len(clean)
		}
	}

	var result strings.Builder
	for i, line := range lines {
		if i > 0 {
			result.WriteString("\n")
		}

		clean := stripAnsi(line)
		switch justify {
		case "center":
			pad := (maxWidth - len(clean)) / 2
			result.WriteString(strings.Repeat(" ", pad) + line)
		case "right":
			pad := maxWidth - len(clean)
			result.WriteString(strings.Repeat(" ", pad) + line)
		default:
			result.WriteString(line)
		}
	}

	return result.String()
}

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

func RenderASCII(input string, bannerLines []string, color, justify string) (string, error) {
	if len(bannerLines) == 0 {
		return "", fmt.Errorf("empty banner file")
	}

	text := strings.ReplaceAll(input, "\\n", "\n")
	textLines := strings.Split(text, "\n")

	var output strings.Builder

	for lineIdx, line := range textLines {
		if line == "" {
			if lineIdx > 0 {
				output.WriteString("\n")
			}
			continue
		}

		for row := 0; row < lineHeight; row++ {
			for _, ch := range line {
				if ch < 32 || ch > 126 {
					continue
				}

				idx := (int(ch)-32)*charBlockSize + 1
				if idx+row >= len(bannerLines) {
					return "", fmt.Errorf("banner file incomplete at character %q", ch)
				}
				output.WriteString(bannerLines[idx+row])
			}
			output.WriteString("\n")
		}
	}

	colored := applyColor(output.String(), color)
	return justifyASCII(colored, justify), nil
}
