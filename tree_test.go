package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildTree(t *testing.T) {
	// Create a temporary directory structure for testing
	tmpDir, err := os.MkdirTemp("", "tree-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test directory structure
	testDirs := []string{
		filepath.Join(tmpDir, "dir1"),
		filepath.Join(tmpDir, "dir1", "subdir1"),
		filepath.Join(tmpDir, "dir2"),
	}

	testFiles := []string{
		filepath.Join(tmpDir, "file1.txt"),
		filepath.Join(tmpDir, "dir1", "file2.txt"),
		filepath.Join(tmpDir, "dir2", "file3.go"),
	}

	// Create directories
	for _, dir := range testDirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
	}

	// Create files
	for _, file := range testFiles {
		if err := os.WriteFile(file, []byte("test"), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	// Test buildTree function
	root, err := buildTree(tmpDir)
	if err != nil {
		t.Fatalf("buildTree failed: %v", err)
	}

	// Verify root node
	if !root.isDir {
		t.Error("Root should be a directory")
	}

	// Test directory counting
	dirCount := countDirs(root)
	expectedDirCount := len(testDirs) + 1 // +1 for root dir
	if dirCount != expectedDirCount {
		t.Errorf("Expected %d directories, got %d", expectedDirCount, dirCount)
	}

	// Test filtering options
	tests := []struct {
		name     string
		opts     *Options
		expected bool
	}{
		{
			name:     "Dirs only",
			opts:     &Options{dirsOnly: true},
			expected: true,
		},
		{
			name:     "Files only",
			opts:     &Options{filesOnly: true},
			expected: false,
		},
		{
			name:     "Pattern match",
			opts:     &Options{pattern: ".go"},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a buffer to capture output
			printTree(root, "", true, tc.opts, 0)
		})
	}
}

func TestTreeNodeStructure(t *testing.T) {
	// Test creating a TreeNode
	node := &TreeNode{
		path:  "test",
		isDir: true,
	}

	// Test adding children
	child1 := &TreeNode{path: "child1", isDir: false}
	child2 := &TreeNode{path: "child2", isDir: true}
	node.children = append(node.children, child1, child2)

	// Verify node structure
	if len(node.children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(node.children))
	}

	if node.children[0].path != "child1" || node.children[0].isDir {
		t.Error("First child node properties are incorrect")
	}

	if node.children[1].path != "child2" || !node.children[1].isDir {
		t.Error("Second child node properties are incorrect")
	}
}
