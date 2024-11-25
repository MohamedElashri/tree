# Poor Man's Tree Command

A lightweight bash implementation of the `tree` command that displays directory structure with colorization and depth control. This script provides basic tree-like directory listing with colored output for different file types (directories, files, symlinks, executables).

## Features

- 🎨 Color-coded output:
  - Directories: Bold Blue
  - Regular Files: Default Color
  - Symbolic Links: Bold Cyan
  - Executable Files: Bold Green
- 📊 Depth control (default: 2 levels)
- 📁 Directory and file count summary
- 🔍 Optional hidden file display
- 📝 Proper handling of symbolic links

## Installation

1. **Create a local bin directory** (if it doesn't exist):
   ```bash
   mkdir -p ~/.local/bin
   ```

2. **Download/Copy the script**:
   ```bash
   # Copy the script to your local bin
   cp tree ~/.local/bin/tree
   ```

3. **Make it executable**:
   ```bash
   chmod +x ~/.local/bin/tree
   ```

4. **Add to PATH** (if not already done):
   Add this line to your `~/.bashrc` or `~/.zshrc`:
   ```bash
   export PATH="$HOME/.local/bin:$PATH"
   ```
   Then reload your shell configuration:
   ```bash
   source ~/.bashrc  # or source ~/.zshrc
   ```

## Usage

```bash
tree [OPTIONS] [directory]
```

### Options

- `-L level`  : Set maximum display depth (default: 2)
- `-a`        : Show all files, including hidden ones
- `-h`        : Display help message

### Examples

```bash
# Show current directory (2 levels deep by default)
tree

# Show specific directory
tree /path/to/directory

# Show 3 levels deep
tree -L 3

# Show unlimited levels
tree -L -1

# Show hidden files
tree -a

# Combine options together
tre -a -L 3 /path/to/directory
```

## Differences from Original Tree Command

This implementation focuses on core functionality and differs from the original `tree` command in several ways:

1. Simplified option set (only -L, -a, and -h are supported)
2. Default depth of 2 levels (matches original tree)
3. Basic color scheme focusing on main file types
4. Simplified output formatting

## Limitations

- Designed for Linux systems (No BSD/MacOS)
- Does not support all options available in the original tree command
- May not handle all special characters in filenames
- No pattern matching or file filtering
- No JSON/XML/HTML output formats

## Troubleshooting

1. **Command not found**:
   - Ensure the script is in your `~/.local/bin`
   - Verify your PATH includes `~/.local/bin`
   - Check if the script is executable

2. **Permission denied**:
   - Check if you have read permissions for the directories
   - Verify the script has execute permissions

3. **No colors showing**:
   - Ensure your terminal supports ANSI colors
   - Check if your terminal's color support is enabled
