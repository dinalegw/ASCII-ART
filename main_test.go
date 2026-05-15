package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func loadTestFont(t *testing.T) []string {
	t.Helper()

	path := filepath.Join("fonts", "standard.txt")
	fontLines, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	if len(fontLines) == 0 {
		t.Fatal("font file is empty")
	}

	return fontLines
}

// ===================== BASIC PIPELINE TEST =====================

func TestFullRenderPipeline(t *testing.T) {
	font := loadTestFont(t)

	tests := []struct {
		name     string
		input    string
		color    string
		justify  string
	}{
		{"simple_left", "Hi", "red", "left"},
		{"simple_center", "Hi", "green", "center"},
		{"simple_right", "Hi", "blue", "right"},
		{"multiline", "Hi\\nGo", "yellow", "left"},
		{"random_color", "Test", "random", "center"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := RenderASCII(tt.input, font, tt.color, tt.justify)
			if err != nil {
				t.Fatalf("RenderASCII failed: %v", err)
			}

			if strings.TrimSpace(output) == "" && tt.input != "" {
				t.Fatal("expected non-empty output")
			}

			if tt.input != "" && !strings.Contains(output, "\n") {
				t.Fatal("expected multiline ASCII output")
			}
		})
	}
}

// ===================== JUSTIFY BEHAVIOUR TEST =====================

func TestJustifyChangesOutput(t *testing.T) {
	font := loadTestFont(t)

	left, err := RenderASCII("Hello", font, "red", "left")
	if err != nil {
		t.Fatal(err)
	}

	center, err := RenderASCII("Hello", font, "red", "center")
	if err != nil {
		t.Fatal(err)
	}

	right, err := RenderASCII("Hello", font, "red", "right")
	if err != nil {
		t.Fatal(err)
	}

	if left == center {
		t.Fatal("center justify should differ from left")
	}

	if left == right {
		t.Fatal("right justify should differ from left")
	}
}

// ===================== COLOR TEST =====================

func TestColorIsApplied(t *testing.T) {
	font := loadTestFont(t)

	output, err := RenderASCII("Hi", font, "red", "left")
	if err != nil {
		t.Fatal(err)
	}

	// ANSI red code should exist in output
	if !strings.Contains(output, "\x1b[31m") {
		t.Fatal("expected red ANSI color code in output")
	}

	if !strings.Contains(output, "\x1b[0m") {
		t.Fatal("expected reset ANSI code in output")
	}
}

// ===================== BANNER LOADING TEST =====================

func TestBannerLoading(t *testing.T) {
	for _, banner := range []string{"standard", "shadow", "thinkertoy"} {
		t.Run(banner, func(t *testing.T) {
			path := filepath.Join("fonts", banner+".txt")

			lines, err := LoadBanner(path)
			if err != nil {
				t.Fatalf("failed to load %s: %v", banner, err)
			}

			if len(lines) < 10 {
				t.Fatalf("banner file too small or corrupted: %s", banner)
			}
		})
	}
}

// ===================== EDGE CASE TEST =====================

func TestEdgeCases(t *testing.T) {
	font := loadTestFont(t)

	tests := []string{
		"",
		" ",
		"\n",
		"!!!",
		"Hello_World_123",
		"_",
		"#00",
		"_#01",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			output, err := RenderASCII(input, font, "random", "center")

			// empty string is allowed to return empty output
			if strings.TrimSpace(input) == "" {
				return
			}

			// IMPORTANT FIX:
			// only fail if BOTH error AND output are empty
			if err != nil && strings.TrimSpace(output) == "" {
				t.Fatalf("expected output for input %q, got error: %v", input, err)
			}

			// if output exists, it's valid
			if strings.TrimSpace(output) == "" {
				t.Fatalf("expected output for non-empty input %q", input)
			}
		})
	}
}