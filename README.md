# Tree Command

A modern implementation of the classic `tree` command in Go, providing a visual directory structure display in your terminal. That's just a poor's man implementation of the GNU utility `tree` because some remote servers doesn't have it. I see you CERN trigger development test servers. 

## Installation

### Prerequisites

- Go 1.21 or higher

### From Source

1. Clone the repository:
```bash
git clone https://github.com/yourusername/tree.git
cd tree
```

2. Build and install:

#### System-wide installation (requires sudo):
```bash
make install
# This will install to /usr/local/bin/tree
```

#### User-specific installation:
```bash
make build
mkdir -p ~/.local/bin
cp tree ~/.local/bin/
# Make sure ~/.local/bin is in your PATH
```

Add this to your shell's configuration file (~/.bashrc, ~/.zshrc, etc.) if ~/.local/bin is not in your PATH:
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