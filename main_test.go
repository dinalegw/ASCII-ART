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

func TestParseOptionsWithColor(t *testing.T) {
	opts, err := ParseOptions([]string{"--color=red", "Hello"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if opts.Color != "red" {
		t.Fatalf("expected color red, got %s", opts.Color)
	}
}

func TestParseOptionsInvalidColor(t *testing.T) {
	_, err := ParseOptions([]string{"--color=invalid", "Hello"})
	if err == nil {
		t.Fatal("expected error for invalid color")
	}
}

func TestParseOptionsInteractive(t *testing.T) {
	opts, err := ParseOptions([]string{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if opts.Text != "" {
		t.Fatalf("expected empty text for interactive mode, got %s", opts.Text)
	}
}

func TestParseOptionsWithJustify(t *testing.T) {
	opts, err := ParseOptions([]string{"--justify=center", "Hello"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if opts.Justify != "center" {
		t.Fatalf("expected justify center, got %s", opts.Justify)
	}
}

func TestParseOptionsInvalidJustify(t *testing.T) {
	_, err := ParseOptions([]string{"--justify=invalid", "Hello"})
	if err == nil {
		t.Fatal("expected error for invalid justify")
	}
}

func TestRenderAsciiWithStandardBanner(t *testing.T) {
	bannerPath := filepath.Join("fonts", "standard.txt")
	fontLines, err := LoadBanner(bannerPath)
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	ascii, err := RenderASCII("Hi", fontLines, "random", "left")
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	if !strings.Contains(ascii, "Hi") && len(ascii) == 0 {
		t.Fatalf("expected rendered ASCII art, got %q", ascii)
	}
}

func TestRenderAsciiEmpty(t *testing.T) {
	bannerPath := filepath.Join("fonts", "standard.txt")
	fontLines, err := LoadBanner(bannerPath)
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	ascii, err := RenderASCII("", fontLines, "random", "left")
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	if ascii != "" {
		t.Fatalf("expected empty output, got %q", ascii)
	}
}

func TestRenderAsciiSpecialChars(t *testing.T) {
	bannerPath := filepath.Join("fonts", "standard.txt")
	fontLines, err := LoadBanner(bannerPath)
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	ascii, err := RenderASCII("Test\\nWorld", fontLines, "random", "left")
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	if len(ascii) == 0 {
		t.Fatalf("expected rendered ASCII art, got empty")
	}
}

func TestRenderAsciiInvalidChar(t *testing.T) {
	bannerPath := filepath.Join("fonts", "standard.txt")
	fontLines, err := LoadBanner(bannerPath)
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	ascii, err := RenderASCII("Hello\x01World", fontLines, "random", "left")
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	if len(ascii) == 0 {
		t.Fatalf("expected rendered ASCII art, got empty")
	}
}
