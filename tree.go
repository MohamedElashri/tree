package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ANSI color codes
const (
	colorReset = "\033[0m"
	colorBlue  = "\033[34m"
	colorGreen = "\033[32m"
)

// Options holds the command-line options
type Options struct {
	maxDepth  int
	pattern   string
	dirsOnly  bool
	filesOnly bool
	noColor   bool
	help      bool
}

func printHelp() {
	fmt.Print(`Usage: tree [options] [path]

Options:
  -L, --level <depth>    Max display depth of the directory tree
  -P, --pattern <pattern> List only files that match the pattern
  -d, --dirs-only        List directories only
  -f, --files-only       List files only
  --no-color            Turn off colorized output
  -h, --help            Show this help message

Example:
  tree
  tree /path/to/directory
  tree -L 2 -P "*.go"
  tree --dirs-only
`)
}

// TreeNode represents a file or directory in the tree
type TreeNode struct {
	path     string
	isDir    bool
	children []*TreeNode
}

// printTree prints the tree structure starting from the given node
func printTree(node *TreeNode, prefix string, isLast bool, opts *Options, depth int) {
	// Prepare the current line prefix
	linePrefix := prefix
	if isLast {
		linePrefix += "└── "
	} else {
		linePrefix += "├── "
	}

	// Check max depth
	if opts.maxDepth > 0 && depth > opts.maxDepth {
		return
	}

	// Check if node matches filters
	if (opts.dirsOnly && !node.isDir) || (opts.filesOnly && node.isDir) {
		return
	}

	baseName := filepath.Base(node.path)
	if opts.pattern != "" && !strings.Contains(baseName, opts.pattern) {
		return
	}

	// Print the current node with color if enabled
	nodeName := baseName
	if !opts.noColor {
		if node.isDir {
			nodeName = colorBlue + nodeName + colorReset
		} else {
			nodeName = colorGreen + nodeName + colorReset
		}
	}
	fmt.Println(linePrefix + nodeName)

	// Prepare the prefix for children
	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	// Print children
	for i, child := range node.children {
		printTree(child, childPrefix, i == len(node.children)-1, opts, depth+1)
	}
}

// buildTree builds a tree structure starting from the given path
func buildTree(path string) (*TreeNode, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	node := &TreeNode{
		path:  path,
		isDir: fileInfo.IsDir(),
	}

	if node.isDir {
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		for _, entry := range entries {
			// Skip hidden files and directories
			if strings.HasPrefix(entry.Name(), ".") {
				continue
			}

			childPath := filepath.Join(path, entry.Name())
			child, err := buildTree(childPath)
			if err != nil {
				continue
			}
			node.children = append(node.children, child)
		}
	}

	return node, nil
}

func main() {
	// Parse command line options
	opts := &Options{}
	flag.IntVar(&opts.maxDepth, "L", 0, "Max display depth of the directory tree")
	flag.StringVar(&opts.pattern, "P", "", "List only files that match the pattern")
	flag.BoolVar(&opts.dirsOnly, "d", false, "List directories only")
	flag.BoolVar(&opts.filesOnly, "f", false, "List files only")
	flag.BoolVar(&opts.noColor, "no-color", false, "Turn off colorized output")
	flag.BoolVar(&opts.help, "h", false, "Show help message")
	flag.Parse()

	// Show help if requested
	if opts.help {
		printHelp()
		os.Exit(0)
	}

	// Get the root path from remaining arguments or use current directory
	rootPath := "."
	if args := flag.Args(); len(args) > 0 {
		rootPath = args[0]
	}

	// Build and print the tree
	root, err := buildTree(rootPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Print the root directory name
	fmt.Println(rootPath)

	// Print all children
	for i, child := range root.children {
		printTree(child, "", i == len(root.children)-1, opts, 1)
	}

	// Print summary
	fmt.Println("\n" + fmt.Sprintf("%d directories", countDirs(root)))
}

// countDirs counts the number of directories in the tree
func countDirs(node *TreeNode) int {
	count := 0
	if node.isDir {
		count++
	}

	for _, child := range node.children {
		count += countDirs(child)
	}

	return count
}
