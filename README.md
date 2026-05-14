# AsciiArtForge

AsciiArtForge is a polished, production-ready command-line tool for generating ASCII art from plain text.
Built in Go with a clean CLI, banner-based typography, file output support, and a user-friendly experience, this project is designed to be the best ASCII art generator for developers, content creators, and terminal power users.

## 🚀 What it does

- Converts text into stylized ASCII art
- Supports multiple banner styles: `standard`, `shadow`, and `thinkertoy`
- Handles multi-line input via `\n`
- Optionally writes output to a text file
- Validates CLI input and provides helpful usage guidance
- Includes automated tests for reliability

## 💡 Why this project

AsciiArtForge is built to meet world-class CLI standards:

- intuitive commands for fast usage
- clean code and modular design
- portable banner files for easy customization
- professional README and licensing for GitHub-ready presentation

## 📦 Installation

Make sure you have Go installed (1.22+ recommended):

```bash
go version
```

Then install or run directly:

```bash
go run . "Hello World"
```

To build a standalone binary:

```bash
go build -o ascii-art-forge .
```

## ▶️ Usage

### Basic text output

```bash
ascii-art-forge "Hello World"
```

### Choose a banner style

```bash
ascii-art-forge --banner=shadow "Hello World"
```

### Save output to file

```bash
ascii-art-forge --output=art.txt "Hello World"
```

### Multi-line text

```bash
ascii-art-forge "Hello\nWorld"
```

## 📁 Project structure

```bash
.
├── LICENSE
├── README.md
├── go.mod
├── main.go
├── cli.go
├── renderer.go
├── main_test.go
├── .gitignore
└── fonts/
    ├── standard.txt
    ├── shadow.txt
    └── thinkertoy.txt
```

## 🧪 Testing

Run the full test suite with:

```bash
go test ./...
```

## 📝 License

This project is released under the MIT License. See `LICENSE` for details.

## 🙌 Contributions

Contributions are welcome. Open issues for feature requests, bug reports, or improvements.
