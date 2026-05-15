package main

import (
	"path/filepath"
	"testing"
)

func TestRenderASCII(t *testing.T) {
	bannerPath := filepath.Join("fonts", "standard.txt")
	fontLines, err := LoadBanner(bannerPath)
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	tests := []struct {
		name   string
		input  string
		color  string
		justify string
	}{
		{"simple", "Hi", "red", "left"},
		{"empty", "", "random", "left"},
		{"multiline", "Hi\\nWorld", "blue", "center"},
		{"special chars", "Test!", "green", "right"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := RenderASCII(tt.input, fontLines, tt.color, tt.justify)
			if err != nil {
				t.Fatalf("RenderASCII failed: %v", err)
			}
			if tt.input != "" && output == "" {
				t.Fatal("expected non-empty output")
			}
		})
	}
}

func TestBannerLoading(t *testing.T) {
	for _, banner := range []string{"standard", "shadow", "thinkertoy"} {
		t.Run(banner, func(t *testing.T) {
			path := filepath.Join("fonts", banner+".txt")
			lines, err := LoadBanner(path)
			if err != nil {
				t.Fatalf("failed to load %s: %v", banner, err)
			}
			if len(lines) == 0 {
				t.Fatal("empty banner file")
			}
		})
	}
}
