# AsciiArtForge

[![CI](https://github.com/dinalegw/AsciiArtForge/actions/workflows/ci.yml/badge.svg)](https://github.com/dinalegw/AsciiArtForge/actions/workflows/ci.yml)

A professional command-line ASCII art generator built in Go. Transform text into colored ASCII art with customizable fonts and alignment.

## Features

- **Interactive Mode**: Run without arguments for guided input
- **Multiple Fonts**: `standard`, `shadow`, `thinkertoy` banner styles
- **Color Support**: red, green, yellow, blue, magenta, cyan, or random
- **Text Alignment**: left, center, or right justification
- **File Output**: Automatically saves to `output.txt`

## Installation

```bash
go install github.com/dinalegw/AsciiArtForge@latest
```

Or build from source:

```bash
git clone https://github.com/dinalegw/AsciiArtForge
cd AsciiArtForge
go build -o ascii-art-forge .
```

## Usage

### Interactive Mode (Recommended)

```bash
./ascii-art-forge
```

Prompts for:
1. Text to convert
2. Banner style (standard/shadow/thinkertoy)
3. Color (red/green/yellow/blue/magenta/cyan/random)
4. Text alignment (left/center/right)

Output is saved to `output.txt`.

### Example Output

```
 _____    ______   _______    _______              _   _  
 / ___ \  |  _  \  | |     |  | ) ) )  |    /\     | \ | | 
| |   | | | | | |  | |     |  |/ / /   |   /  \    |  \| | 
| |   | | | |  _|  |  >    |      <    |  / /\ \   | . ` | 
| |___| | | | |    | |     |  |\_  |   | |  __  |   | |\  | 
 \_____/  |_| |    |_|_____|  |___| |  |__\_|_|_|__|_\_| 
```

## Project Structure

```
AsciiArtForge/
├── main.go           # Entry point and CLI
├── renderer.go       # ASCII rendering logic
├── main_test.go      # Test suite
├── fonts/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── LICENSE           # MIT License
└── README.md
```

## Testing

```bash
go test -v ./...
```

## License

MIT License - see [LICENSE](LICENSE) for details.

## Contributing

Contributions welcome! Open an issue or submit a pull request.