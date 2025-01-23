# Tree Command

A modern implementation of the classic `tree` command in Go, providing a visual directory structure display in your terminal. That's just a poor's man implementation of the GNU utility `tree` because some remote servers doesn't have it. I see you CERN trigger development test servers. 

## Installation

### Prerequisites

- Go 1.21 or higher (only needed for building from source)

### Pre-built Binaries

You can download pre-built binaries for Linux and macOS from the [latest release](https://github.com/MohamedElashri/tree/releases/latest).

#### System-wide installation:

For Linux (AMD64)

```bash

curl -L https://github.com/MohamedElashri/tree/releases/latest/download/tree-linux-amd64 -o tree
sudo mv tree /usr/local/bin/
sudo chmod +x /usr/local/bin/tree
```
For macOS (AMD64)

```bash
curl -L https://github.com/MohamedElashri/tree/releases/latest/download/tree-darwin-amd64 -o tree
sudo mv tree /usr/local/bin/
sudo chmod +x /usr/local/bin/tree
```

#### User-specific installation:

For Linux (AMD64)

```bash
curl -L https://github.com/MohamedElashri/tree/releases/latest/download/tree-linux-amd64 -o tree
mkdir -p ~/.local/bin
mv tree ~/.local/bin/
chmod +x ~/.local/bin/tree
```

For macOS (AMD64)
```bash
curl -L https://github.com/MohamedElashri/tree/releases/latest/download/tree-darwin-amd64 -o tree
mkdir -p ~/.local/bin
mv tree ~/.local/bin/
chmod +x ~/.local/bin/tree
```

Make sure `~/.local/bin` is in your PATH (add to ~/.bashrc, ~/.zshrc, etc. if needed):
```bash
export PATH="$HOME/.local/bin:$PATH"
```

### From Source

1. Clone the repository:
```bash
git clone https://github.com/MohamedElashri/tree
cd tree
```

2. Build and install:

#### System-wide installation (requires sudo):

This will install to `/usr/local/bin/tree`

```bash
make install
```

#### User-specific installation:
```bash
make build
mkdir -p ~/.local/bin
cp tree ~/.local/bin/
```

Make sure` ~/.local/bin` is in your PATH


Add this to your shell's configuration file (`~/.bashrc`, `~/.zshrc`, etc.) if` ~/.local/bin` is not in your PATH:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

## Usage

```bash
tree [options] [directory]
```

### Available Commands

| Command | Description |
|---------|-------------|
| `tree` | Display directory structure for current directory |
| `tree <directory>` | Display directory structure for specified directory |
| `tree -L <level>` | Limit display to specified depth level |
| `tree -d` | List directories only |
| `tree -f` | Print full path prefix for each file |
| `tree -i` | Don't print indentation lines |
| `tree -p` | Print file type and permissions |
| `tree -s` | Print size of each file |
| `tree -h` | Print size in human readable format |
| `tree --help` | Display help information |
| `tree -v, --version` | Display current version |
| `tree --update` | Update to the latest version |

## Development

### Available Make Commands

| Command | Description |
|---------|-------------|
| `make` | Build the binary |
| `make build` | Build with optimizations |
| `make clean` | Remove build artifacts |
| `make install` | Install system-wide |
| `make test` | Run tests |
| `make fmt` | Format code |

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.