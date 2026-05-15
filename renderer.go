package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
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

// ================= LOAD BANNER =================

func LoadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

// ================= RENDER ENGINE =================

func RenderASCII(input string, bannerLines []string, color, justify string) (string, error) {
	if len(bannerLines) == 0 {
		return "", fmt.Errorf("empty banner file")
	}

	text := strings.ReplaceAll(input, "\\n", "\n")
	blocks := strings.Split(text, "\n")

	var result strings.Builder

	for _, block := range blocks {
		if strings.TrimSpace(block) == "" {
			result.WriteString("\n")
			continue
		}

		words := strings.Fields(block)

		// build ASCII per word (8 rows each)
		wordArt := make([][]string, len(words))

		for wi, w := range words {
			wordArt[wi] = make([]string, lineHeight)

			for row := 0; row < lineHeight; row++ {
				var sb strings.Builder

				for _, ch := range w {
					idx := (int(ch)-32)*charBlockSize + 1
					sb.WriteString(bannerLines[idx+row])
				}

				wordArt[wi][row] = sb.String()
			}
		}

		// compute ASCII WIDTH (IMPORTANT FIX)
		width := getTerminalWidth()

		asciiWordWidth := 0
		for _, w := range wordArt {
			asciiWordWidth += len(w[0])
		}

		gaps := len(words) - 1
		extraSpace := 0
		baseSpace := 1

		if justify == "justify" && gaps > 0 {
			spacesNeeded := width - asciiWordWidth

			if spacesNeeded > 0 {
				baseSpace = spacesNeeded / gaps
				extraSpace = spacesNeeded % gaps
			}
		}

		// render 8 rows
		for row := 0; row < lineHeight; row++ {
			var line strings.Builder

			for i := 0; i < len(words); i++ {
				line.WriteString(wordArt[i][row])

				if i < gaps {
					if justify == "justify" {
						sp := baseSpace
						if extraSpace > 0 {
							sp++
							extraSpace--
						}
						line.WriteString(strings.Repeat(" ", sp))
					} else {
						line.WriteString(" ")
					}
				}
			}

			result.WriteString(line.String() + "\n")
		}
	}

	return applyColor(result.String(), color), nil
}

// ================= COLOR =================

func applyColor(asciiArt string, color string) string {
	lines := strings.Split(asciiArt, "\n")
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		var code string
		if color == "random" {
			code = randomColors[rand.Intn(len(randomColors))]
		} else {
			code = colorMap[color]
		}

		result.WriteString(code)
		result.WriteString(line)
		result.WriteString(resetCode)
		result.WriteString("\n")
	}

	return result.String()
}

// ================= TERMINAL WIDTH =================

func getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	out, err := cmd.Output()
	if err != nil {
		return 80
	}

	parts := strings.Fields(string(out))
	if len(parts) != 2 {
		return 80
	}

	w, err := strconv.Atoi(parts[1])
	if err != nil {
		return 80
	}

	return w
}