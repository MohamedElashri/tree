# Release Notes

## v0.0.2
This release adds version management capabilities to help users keep their installation up to date.

### Features
- **Version Checking**: Added `-v` or `--version` flag to display current version
- **Auto Update**: Added `--update` flag to automatically update to the latest version
- **Internal Updates**: Implemented version management system using GitHub's release infrastructure

## v0.0.1 (Initial Release)
This is  the first release of `Go` implementation of the classic `tree` GNU utility! This utility provides a clean and efficient way to visualize directory structures in terminal.

### Features

- **Directory Structure Visualization**: Display hierarchical file and directory structures in a clear, tree-like format
- **Flexible Installation Options**: Support for both system-wide and user-specific installations
- **Cross-Platform Support**: Pre-built binaries available for Linux and macOS (AMD64)
- **Customizable Display Options**: Various flags to control the output format and depth

### Installation Methods

#### Pre-built Binaries
- Direct download and installation for Linux and macOS
- Support for both system-wide (`/usr/local/bin`) and user-specific (`~/.local/bin`) installations
- Simple curl-based installation process

#### Source Installation
- Build from source with Go 1.21 or higher
- Makefile support for easy building and installation
- Support for both system-wide and user-specific compilation and installation

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

### Development Tools

The following make commands are available for development:

| Command | Description |
|---------|-------------|
| `make` | Build the binary |
| `make build` | Build with optimizations |
| `make clean` | Remove build artifacts |
| `make install` | Install system-wide |
| `make test` | Run tests |
| `make fmt` | Format code |

### Notes

- This is the initial release focusing on core functionality
- The implementation is designed to be lightweight and efficient
- Compatible with Go 1.21 and higher
- Licensed under MIT License

### Future Plans

- Additional platform support
- More customization options
- Performance optimizations
- Enhanced filtering capabilities