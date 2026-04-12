// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"os"
	"path/filepath"
	"testing"

	"gohan/src/env"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("could not write test file %s: %v", path, err)
	}
}

func TestReadPalette_Success(t *testing.T) {
	dir := t.TempDir()

	file := filepath.Join(dir, "test.palette")
	writeFile(t, file, "palette-content")

	paths := env.Paths{
		PalettesDir: dir,
	}

	result, err := readPalette("test.palette", paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != "palette-content" {
		t.Errorf("expected %q, got %q", "palette-content", result)
	}
}

func TestReadPalette_NotFound(t *testing.T) {
	paths := env.Paths{
		PalettesDir: t.TempDir(),
	}

	_, err := readPalette("missing.palette", paths)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestReadCurrentPalette_Success(t *testing.T) {
	file := t.TempDir() + "/current.palette"

	writeFile(t, file, "current-content")

	paths := env.Paths{
		CurrentPaletteFile: file,
	}

	result, err := readCurrentPalette(paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != "current-content" {
		t.Errorf("expected %q, got %q", "current-content", result)
	}
}

func TestReadCurrentPalette_NotFound(t *testing.T) {
	paths := env.Paths{
		CurrentPaletteFile: "/non/existent/file",
	}

	_, err := readCurrentPalette(paths)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
