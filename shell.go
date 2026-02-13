package system

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// ShellCommandResult holds the output and exit code of a shell command.
type ShellCommandResult struct {
	StdOut   string
	StdErr   string
	ExitCode int
	Args     []string
}

// HasError returns true if stderr contains output.
func (r ShellCommandResult) HasError() bool {
	return len(r.StdErr) > 0
}

// GetError returns stderr content as an error.
func (r ShellCommandResult) GetError() error {
	return errors.New(r.StdErr)
}

// Exec runs a shell command and returns the captured stdout, stderr, and exit code.
func Exec(args ...string) (*ShellCommandResult, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		return nil, err
	}

	c := ShellCommandResult{
		Args:     args,
		StdOut:   stdout.String(),
		StdErr:   stderr.String(),
		ExitCode: cmd.ProcessState.ExitCode(),
	}

	if c.HasError() {
		return &c, c.GetError()
	}

	return &c, nil
}

// FileInfo holds parsed components of a file path.
type FileInfo struct {
	Name     string // Full file name with extension
	BaseName string // File name without extension
	Dir      string // Directory path
	Ext      string // File extension including dot
}

// GetFileInfo splits a file path into its components.
func GetFileInfo(path string) FileInfo {
	dir, fileFullName := filepath.Split(path)
	ext := filepath.Ext(path)
	fileName := fileFullName[0 : len(fileFullName)-len(ext)]

	return FileInfo{
		Name:     fileFullName,
		BaseName: fileName,
		Dir:      dir,
		Ext:      ext,
	}
}

// IsFileExists checks if a file exists and is not a directory.
func IsFileExists(f string) bool {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// IsLaunchedByDebugger checks if the current process was launched by the Delve debugger.
// Requires gops (https://github.com/google/gops) to be in PATH.
func IsLaunchedByDebugger() bool {
	gopsOut, err := exec.Command("gops", strconv.Itoa(os.Getppid())).Output()
	if err == nil && strings.Contains(string(gopsOut), "\\dlv.exe") {
		return true
	}
	return false
}
