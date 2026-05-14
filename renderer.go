package main

import (
	"fmt"
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

// RenderASCII converts input text into ASCII art using the provided banner.
func RenderASCII(input string, bannerLines []string) (string, error) {
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

	return output.String(), nil
}
