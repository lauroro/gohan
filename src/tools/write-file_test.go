// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package tools

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFile_Success(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "test.txt")

	err := WriteFile(file, "hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	if string(data) != "hello" {
		t.Errorf("expected %q, got %q", "hello", string(data))
	}
}

func TestWriteFile_Overwrite(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "test.txt")

	if err := os.WriteFile(file, []byte("old"), 0644); err != nil {
		t.Fatal(err)
	}

	err := WriteFile(file, "new")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, _ := os.ReadFile(file)

	if string(data) != "new" {
		t.Errorf("expected overwrite, got %q", string(data))
	}
}

func TestWriteFile_CreatesDirectories(t *testing.T) {
	dir := t.TempDir()

	file := filepath.Join(dir, "nested", "dir", "file.txt")

	err := WriteFile(file, "hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	if string(data) != "hello" {
		t.Errorf("expected %q, got %q", "hello", string(data))
	}
}

func TestWriteFile_InvalidPath(t *testing.T) {
	err := WriteFile("/invalid/path/test.txt", "data")
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestWriteFile_EmptyContent(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "empty.txt")

	err := WriteFile(file, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, _ := os.ReadFile(file)

	if string(data) != "" {
		t.Errorf("expected empty file, got %q", string(data))
	}
}
