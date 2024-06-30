package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// CheckError now returns an error instead of exiting.
func CheckError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %v", message, err)
	}
}

// AppendToFile is updated to use os.WriteFile for appending content to a file.
func AppendToFile(filename, content string) {
	// Open or create the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		CheckError(err, fmt.Sprintf("Failed to open file: %s", filename))
	}
	defer f.Close()

	// Write the content
	_, err = f.WriteString(content)
	CheckError(err, fmt.Sprintf("Failed to write to file: %s", filename))
}

// // RunCommand now returns output to be more flexible.
// func RunCommand(command string, args ...string) (string, error) {
// 	cmd := exec.Command(command, args...)
// 	output, err := cmd.CombinedOutput()
// 	return string(output), err
// }

// RunCommand executes a shell command and streams its output.
func RunCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// RunCommandWithOutput executes a shell command and returns its output.
func RunCommandWithOutput(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// InstallPackage simplified error handling.
func InstallPackage(pkg string) {
	fmt.Printf("Installing %s...\n", pkg)
	err := RunCommand("brew", "install", pkg)
	CheckError(err, fmt.Sprintf("Failed to install %s", pkg))
}

// DownloadAndExtract using Go's http and zip packages.
func DownloadAndExtract(url, destDir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download from %s: %v", url, err)
	}
	defer resp.Body.Close()

	tempFile, err := os.CreateTemp("", "temp-*.zip")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to temp file: %v", err)
	}

	zipReader, err := zip.OpenReader(tempFile.Name())
	if err != nil {
		return fmt.Errorf("failed to open zip file: %v", err)
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		// Construct the full path for the file/directory
		fullPath := filepath.Join(destDir, file.Name)

		// Check if the file is a directory
		if file.FileInfo().IsDir() {
			// Create the directory
			os.MkdirAll(fullPath, os.ModePerm)
			continue
		}

		// Create the directory structure for the file
		if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory for %s: %v", fullPath, err)
		}

		// Open the file within the zip archive
		inFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open zip file %s: %v", file.Name, err)
		}
		defer inFile.Close()

		// Create a new file in the destination directory
		outFile, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return fmt.Errorf("failed to open destination file %s: %v", fullPath, err)
		}

		// Copy the file contents from the zip archive to the new file
		if _, err = io.Copy(outFile, inFile); err != nil {
			outFile.Close() // Close the file explicitly on error
			return fmt.Errorf("failed to copy contents to %s: %v", fullPath, err)
		}

		// Close the new file
		outFile.Close()
	}

	return nil
}

// GetDefaultShell returns the default shell path.
func GetDefaultShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/bash"
	}
	return shell
}

// IsCommandAvailable checks if a command is available in the system.
func IsCommandAvailable(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// IsPackageInstalled checks if a package is installed using brew.
func IsPackageInstalled(pkg string) bool {
	output, err := RunCommandWithOutput("brew", "list", "--formula", pkg)
	if err != nil {
		return false
	}

	return strings.Contains(output, pkg)
}

// IsGemInstalled checks if a gem is installed.
func IsGemInstalled(gem string) bool {
	output, err := RunCommandWithOutput("gem", "list", gem)
	if err != nil {
		return false
	}

	return strings.Contains(output, gem)
}

// AppendLinesToZshrc appends lines to the .zshrc file.
func AppendLinesToZshrc(lines ...string) {
	homeDir, err := os.UserHomeDir()
	CheckError(err, "Failed to get home directory")

	zshrc := filepath.Join(homeDir, ".zshrc")
	for _, line := range lines {
		AppendToFile(zshrc, line+"\n")
	}
}
