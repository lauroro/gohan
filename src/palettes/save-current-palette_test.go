// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"os"
	"path/filepath"
	"testing"

	"codeberg.org/roro/gohan/src/env"
)

func TestSaveCurrentPalette_Success(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "current.yaml")

	paths := env.Paths{
		CurrentPaletteFile: file,
	}

	err := SaveCurrentPalette("content", paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	if string(data) != "content" {
		t.Errorf("expected %q, got %q", "content", string(data))
	}
}

func TestSaveCurrentPalette_Overwrite(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "current.yaml")

	if err := os.WriteFile(file, []byte("old"), 0644); err != nil {
		t.Fatal(err)
	}

	paths := env.Paths{
		CurrentPaletteFile: file,
	}

	err := SaveCurrentPalette("new", paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, _ := os.ReadFile(file)

	if string(data) != "new" {
		t.Errorf("expected overwrite, got %q", string(data))
	}
}

func TestSaveCurrentPalette_CreatesDirectories(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "nested", "dir", "current.yaml")

	paths := env.Paths{
		CurrentPaletteFile: file,
	}

	err := SaveCurrentPalette("content", paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	if string(data) != "content" {
		t.Errorf("expected %q, got %q", "content", string(data))
	}
}

func TestSaveCurrentPalette_EmptyContent(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "current.yaml")

	paths := env.Paths{
		CurrentPaletteFile: file,
	}

	err := SaveCurrentPalette("", paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, _ := os.ReadFile(file)

	if string(data) != "" {
		t.Errorf("expected empty file, got %q", string(data))
	}
}

func TestSaveCurrentPalette_InvalidPath(t *testing.T) {
	paths := env.Paths{
		CurrentPaletteFile: "/invalid/path/current.yaml",
	}

	err := SaveCurrentPalette("content", paths)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
