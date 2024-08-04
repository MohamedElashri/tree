# Poor man tree command 

This script mimics the basic functionality of the `tree` command with a focus on the `-L` option, allowing you to list directory contents up to a specified depth.

## Features

- List directory contents in a tree-like format.
- Specify the maximum depth for listing with a default of 1.
- Defaults to the current directory if no path is provided.
- Simple help option to understand usage.

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

- **List a Specific Directory**: Specify a directory to list its contents.

  ```bash
  tree /path/to/directory
  ```

- **Specify a Depth**: Use a second argument to set the maximum depth.

  ```bash
  tree /path/to/directory 3
  ```

- **Help**: View the usage instructions.

  ```bash
  tree -h
  ```

## Notes

- Ensure your terminal has the necessary permissions to read the directories you want to list.
- This script provides basic functionality and may not cover all edge cases of the full `tree` command.
- Customize the script as needed to fit your specific requirements.

## Troubleshooting

- If the command is not found, ensure that `$HOME/.local/bin` is in your `PATH`.
- If you encounter permission issues, check the directory permissions you are attempting to list.
```

### Key Points:
- **Installation**: Steps guide you through creating a `bin` directory, downloading the script, and setting up the environment.
- **Usage**: Provides examples of how to use the script from any location.
- **Customization**: Encourages modifying the script for specific needs.

With these instructions, you can easily install and use your custom `tree` script from anywhere on your system.
