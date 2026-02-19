package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExpandTilde_withTilde(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("failed to get home dir: %v", err)
	}

	got := ExpandTilde("~/some/path")
	want := filepath.Join(home, "some/path")
	if got != want {
		t.Errorf("ExpandTilde(\"~/some/path\") = %q, want %q", got, want)
	}
}

func TestExpandTilde_withoutTilde(t *testing.T) {
	got := ExpandTilde("/absolute/path")
	want := "/absolute/path"
	if got != want {
		t.Errorf("ExpandTilde(\"/absolute/path\") = %q, want %q", got, want)
	}
}

func TestExpandTilde_relativePath(t *testing.T) {
	got := ExpandTilde("relative/path")
	want := "relative/path"
	if got != want {
		t.Errorf("ExpandTilde(\"relative/path\") = %q, want %q", got, want)
	}
}

func TestDefaultConfigDir(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("failed to get home dir: %v", err)
	}

	got := DefaultConfigDir()
	want := filepath.Join(home, ".config/teams")
	if got != want {
		t.Errorf("DefaultConfigDir() = %q, want %q", got, want)
	}
}

func TestDefaultLayoutPath(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("failed to get home dir: %v", err)
	}

	got := DefaultLayoutPath()
	want := filepath.Join(home, ".config/teams/teams.kdl")
	if got != want {
		t.Errorf("DefaultLayoutPath() = %q, want %q", got, want)
	}
}
