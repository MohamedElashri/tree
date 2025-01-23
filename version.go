package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	version   = "0.0.2"
	githubAPI = "https://api.github.com/repos/MohamedElashri/tree/releases/latest"
)

type GithubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func printVersion() {
	fmt.Printf("tree version %s\n", version)
}

func checkForUpdate() error {
	// Get latest release info from GitHub
	resp, err := http.Get(githubAPI)
	if err != nil {
		return fmt.Errorf("failed to check for updates: %v", err)
	}
	defer resp.Body.Close()

	var release GithubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return fmt.Errorf("failed to parse release info: %v", err)
	}

	// Compare versions
	latestVersion := strings.TrimPrefix(release.TagName, "v")
	if latestVersion == version {
		fmt.Println("You are using the latest version.")
		return nil
	}

	fmt.Printf("New version %s is available (current: %s)\n", latestVersion, version)
	return nil
}

func updateTool() error {
	// Get latest release info
	resp, err := http.Get(githubAPI)
	if err != nil {
		return fmt.Errorf("failed to get release info: %v", err)
	}
	defer resp.Body.Close()

	var release GithubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return fmt.Errorf("failed to parse release info: %v", err)
	}

	// Find the appropriate asset for current platform
	platform := fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
	var downloadURL string
	for _, asset := range release.Assets {
		if strings.Contains(asset.Name, platform) {
			downloadURL = asset.BrowserDownloadURL
			break
		}
	}

	if downloadURL == "" {
		return fmt.Errorf("no compatible binary found for your platform")
	}

	// Download the new binary
	resp, err = http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download update: %v", err)
	}
	defer resp.Body.Close()

	// Create temporary file
	tmp, err := os.CreateTemp("", "tree-*")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tmp.Name())

	// Copy downloaded content to temporary file
	if _, err := io.Copy(tmp, resp.Body); err != nil {
		return fmt.Errorf("failed to save update: %v", err)
	}
	tmp.Close()

	// Make temporary file executable
	if err := os.Chmod(tmp.Name(), 0755); err != nil {
		return fmt.Errorf("failed to make binary executable: %v", err)
	}

	// Get current executable path
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get current executable path: %v", err)
	}
	exePath, err = filepath.EvalSymlinks(exePath)
	if err != nil {
		return fmt.Errorf("failed to resolve symlinks: %v", err)
	}

	// Replace current binary with new one
	if err := os.Rename(tmp.Name(), exePath); err != nil {
		return fmt.Errorf("failed to install update: %v", err)
	}

	fmt.Println("Successfully updated to latest version!")
	return nil
}
