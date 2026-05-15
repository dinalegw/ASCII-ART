# AsciiArtForge

[![CI](https://github.com/dinalegw/AsciiArtForge/actions/workflows/ci.yml/badge.svg)](https://github.com/dinalegw/AsciiArtForge/actions/workflows/ci.yml)

A professional command-line ASCII art generator built in Go. Transform text into colored ASCII art with customizable fonts and alignment.

## Demo

```
  ____    ____    _____    _____              _   _  
 / __ \  |  _ \  |  __ \  |_   _|     /\     | \ | | 
| |  | | | |_) | | |__) |   | |      /  \    |  \| | 
| |  | | |  _ <  |  _  /    | |     / /\ \   | . ` | 
| |__| | | |_) | | | \ \   _| |_   / ____ \  | |\  | 
 \____/  |____/  |_|  \_\ |_____| /_/    \_\ |_| \_| 
                                                    
                                                    

              _   _       _ _       
             | | | |     | | |      
             | |_| | __ _| | | ___  
             |  _  |/ _` | | |/ _ \ 
             | | | | (_| | | |  __/ 
             \_| |_/\__,_|_|_|\___| 
                                   
```

Rendered in red ANSI color → saved to `output.txt`

---

## Features

- **Interactive Console Mode** - Guided prompts for text, banner, color, and alignment
- **Multiple Banner Styles** - `standard`, `shadow`, `thinkertoy` 
- **Color Output** - red, green, yellow, blue, magenta, cyan, or random per character
- **Text Alignment** - left, center, right justification
- **File Output** - Automatically saves to `output.txt`

---

## Installation

### From Source

```bash
git clone https://github.com/dinalegw/AsciiArtForge
cd AsciiArtForge
go build -o ascii-art-forge .
```

### Requirements

- Go 1.22 or higher

---

## Usage

### Interactive Mode (Recommended)

Run the binary without arguments for an interactive prompt:

```bash
./ascii-art-forge
```

You will be prompted to enter:

1. **Text** - The text to convert to ASCII art
2. **Banner** - Font style (`standard`, `shadow`, or `thinkertoy`)
3. **Color** - ANSI color (`red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `random`)
4. **Justify** - Text alignment (`left`, `center`, `right`)

The result is saved to `output.txt`.

### Example Session

```
$ ./ascii-art-forge
Enter text: Hello World
Banner (standard/shadow/thinkertoy) [standard]: shadow
Color (red/green/yellow/blue/magenta/cyan/random) [random]: red
Justify (left/center/right) [left]: center
ASCII art saved to output.txt
```

---

## Project Structure

```
AsciiArtForge/
├── main.go           # Entry point and interactive CLI
├── renderer.go       # ASCII rendering logic
├── main_test.go      # Test suite
├── README.md         # This file
├── LICENSE           # MIT License
├── go.mod            # Go module definition
└── fonts/
    ├── standard.txt  # Standard ASCII font
    ├── shadow.txt    # Shadow style font
    └── thinkertoy.txt # Thinkertoy style font
```

---

## Testing

```bash
go test -v ./...
```

---

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) for details.

---

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.