package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestParseOptionsDefault(t *testing.T) {
	opts, err := ParseOptions([]string{"Hello, World!"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if opts.Banner != "standard" {
		t.Fatalf("expected default banner standard, got %s", opts.Banner)
	}
	if opts.Text != "Hello, World!" {
		t.Fatalf("expected text preserved, got %s", opts.Text)
	}
}

func TestParseOptionsWithOutputAndBanner(t *testing.T) {
	opts, err := ParseOptions([]string{"--output=art.txt", "Hello", "shadow"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if opts.OutputFile != "art.txt" {
		t.Fatalf("expected output art.txt, got %s", opts.OutputFile)
	}
	if opts.Banner != "shadow" {
		t.Fatalf("expected banner shadow, got %s", opts.Banner)
	}
}

func TestParseOptionsInvalidBanner(t *testing.T) {
	_, err := ParseOptions([]string{"--banner=invalid", "Hello"})
	if err == nil {
		t.Fatal("expected error for invalid banner")
	}
}

func TestRenderAsciiWithStandardBanner(t *testing.T) {
	bannerPath := filepath.Join("fonts", "standard.txt")
	fontLines, err := LoadBanner(bannerPath)
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	ascii, err := RenderASCII("Hi", fontLines, false)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	if !strings.Contains(ascii, "Hi") && len(ascii) == 0 {
		t.Fatalf("expected rendered ASCII art, got %q", ascii)
	}
}
