#!/bin/bash

# Function to list directories up to a specified depth
list_directory() {
    local path="$1"
    local level="$2"
    local max_depth="$3"
    local prefix="$4"

    # Stop if current level exceeds max depth
    if [[ "$level" -gt "$max_depth" ]]; then
        return
    fi

    # Get the list of entries in the directory
    local entries
    entries=$(ls -A "$path" 2>/dev/null)

    # Loop through the entries
    local last_index=$(($(echo "$entries" | wc -l) - 1))
    local index=0
    while IFS= read -r entry; do
        local entry_path="$path/$entry"
        local connector new_prefix

        if [[ "$index" -eq "$last_index" ]]; then
            connector="└── "
            new_prefix="$prefix    "
        else
            connector="├── "
            new_prefix="$prefix│   "
        fi

        # Print the entry
        echo "${prefix}${connector}${entry}"

        # If the entry is a directory, recurse into it
        if [[ -d "$entry_path" ]]; then
            list_directory "$entry_path" $((level + 1)) "$max_depth" "$new_prefix"
        fi

        index=$((index + 1))
    done <<< "$entries"
}

# Main function to initiate the tree listing
tree() {
    local path="${1:-.}"     # Use current directory if no path is specified
    local max_depth="${2:-1}" # Default depth is 1

    # Print the initial directory
    echo "$path"

    # Start listing from the initial path
    list_directory "$path" 1 "$max_depth" ""
}

# Check for help argument
if [[ "$1" == "-h" || "$1" == "--help" ]]; then
    echo "Usage: $0 [directory] [max_depth]"
    echo "If no directory is specified, the current directory is used."
    echo "If no max_depth is specified, the default is 1."
    exit 0
fi

# Run the tree function with optional arguments
tree "$1" "$2"
