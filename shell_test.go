package system

import (
	"testing"
)

func TestExec(t *testing.T) {
	result, err := Exec("echo", "hello")
	if err != nil {
		t.Fatalf("Exec() error = %v", err)
	}
	if result.StdOut != "hello\n" {
		t.Errorf("Exec() StdOut = %q, want %q", result.StdOut, "hello\n")
	}
	if result.ExitCode != 0 {
		t.Errorf("Exec() ExitCode = %d, want 0", result.ExitCode)
	}
}

func TestExecError(t *testing.T) {
	_, err := Exec("nonexistent-command-xyz")
	if err == nil {
		t.Error("Exec() expected error for nonexistent command")
	}
}

func TestGetFileInfo(t *testing.T) {
	info := GetFileInfo("/home/user/documents/report.pdf")
	if info.Name != "report.pdf" {
		t.Errorf("Name = %s, want report.pdf", info.Name)
	}
	if info.BaseName != "report" {
		t.Errorf("BaseName = %s, want report", info.BaseName)
	}
	if info.Ext != ".pdf" {
		t.Errorf("Ext = %s, want .pdf", info.Ext)
	}
	if info.Dir != "/home/user/documents/" {
		t.Errorf("Dir = %s, want /home/user/documents/", info.Dir)
	}
}

func TestIsFileExists(t *testing.T) {
	if !IsFileExists("/etc/hosts") {
		t.Error("IsFileExists(/etc/hosts) = false, want true")
	}
	if IsFileExists("/nonexistent/file") {
		t.Error("IsFileExists returned true for nonexistent file")
	}
}
