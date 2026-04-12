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
		t.Fatal(err)
	}
}

func TestLoadPalette_Success(t *testing.T) {
	dir := t.TempDir()

	content := `
base00: "#000000"
base01: "#111111"
base02: "#222222"
base03: "#333333"
base04: "#444444"
base05: "#555555"
base06: "#666666"
base07: "#777777"
base08: "#888888"
base09: "#999999"
base0A: "#aaaaaa"
base0B: "#bbbbbb"
base0C: "#cccccc"
base0D: "#dddddd"
base0E: "#eeeeee"
base0F: "#ffffff"
`

	file := filepath.Join(dir, "dracula.yaml")
	writeFile(t, file, content)

	paths := env.Paths{
		PalettesDir: dir,
	}

	p, err := LoadPalette("dracula", paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p.Base00 != "000000" {
		t.Errorf("expected normalized value, got %s", p.Base00)
	}
}

func TestLoadPalette_FileMissing(t *testing.T) {
	paths := env.Paths{
		PalettesDir: t.TempDir(),
	}

	_, err := LoadPalette("missing", paths)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestLoadPalette_InvalidPalette(t *testing.T) {
	dir := t.TempDir()

	file := filepath.Join(dir, "bad.yaml")
	writeFile(t, file, `base00: "zzzzzz"`)

	paths := env.Paths{
		PalettesDir: dir,
	}

	_, err := LoadPalette("bad", paths)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
