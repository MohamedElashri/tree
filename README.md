# Poor Man Tree Command

This script mimics the basic functionality of the `tree` command, allowing you to list directory contents in a tree-like format with a focus on specifying depth.

## Installation

Follow these steps to install the script and make it available globally on your system:

1. **Create a `bin` Directory**: Ensure you have a `bin` directory in your home folder.

   ```bash
   mkdir -p $HOME/.local/bin
   ```

2. **Download the Script**: Create a new file for the script, or download it directly.

   ```bash
   curl -o $HOME/.local/bin/tree https://github.com/MohamedElashri/tree/raw/main/tree
   ```

   Or copy the script contents manually into a new file:

   ```bash
   nano $HOME/.local/bin/tree
   ```

   Paste the script and save the file.

3. **Make the Script Executable**: Change the permissions of the script to make it executable.

   ```bash
   chmod +x $HOME/.local/bin/tree
   ```

4. **Add to Path**: Ensure your `.bashrc` or `.zshrc` includes `$HOME/.local/bin` in your `PATH`. Add the following line if it isn't present:

   ```bash
   export PATH=$HOME/.local/bin:$PATH
   ```

   Then, source your shell configuration to update the `PATH`:

   ```bash
   source ~/.bashrc
   # or
   source ~/.zshrc
   ```

## Usage

You can use the script from any terminal session. Here's how to use it:

- **List the Current Directory**: Run the script without arguments to list the current directory up to a depth of 1.

  ```bash
  tree
  ```

- **List a Specific Directory**: Use the `-p` or `--path` option to specify a directory to list its contents.

  ```bash
  tree -p /path/to/directory
  ```

- **Specify a Depth**: Use the `-d` or `--depth` option to set the maximum depth.

  ```bash
  tree -p /path/to/directory -d 3
  ```

- **List the Current Directory with a Specific Depth**: Specify only the depth for the current directory.

  ```bash
  tree -d 3
  ```

- **Help**: View the usage instructions.

  ```bash
  tree -h
  ```

## Notes

- Ensure your terminal has the necessary permissions to read the directories you want to list.
- This script provides basic functionality and may not cover all edge cases of the full `tree` command.

## Troubleshooting

- If the command is not found, ensure that `$HOME/.local/bin` is in your `PATH`.
- If you encounter permission issues, check the directory permissions you are attempting to list.
