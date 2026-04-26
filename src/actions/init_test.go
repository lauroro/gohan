// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package actions

import (
	"os"
	"path/filepath"
	"testing"

	"gohan/src/env"
)

func TestInit_Success(t *testing.T) {
	dir := t.TempDir()

	palettesDir := filepath.Join(dir, "palettes")
	os.MkdirAll(palettesDir, 0755)

	content := `base00: "000000"
base01: "111111"
base02: "222222"
base03: "333333"
base04: "444444"
base05: "555555"
base06: "666666"
base07: "777777"
base08: "888888"
base09: "999999"
base0A: "aaaaaa"
base0B: "bbbbbb"
base0C: "cccccc"
base0D: "dddddd"
base0E: "eeeeee"
base0F: "ffffff"
`
	writeFile(t, filepath.Join(palettesDir, "dark.yaml"), content)

	currentPath := filepath.Join(palettesDir, "current.yaml")

	paths := env.Paths{
		PalettesDir:        palettesDir,
		CurrentPaletteFile: currentPath,
	}

	err := Init("dark", paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	current := readFile(t, currentPath)

	if current != content {
		t.Errorf("current palette mismatch\nexpected:\n%s\ngot:\n%s", content, current)
	}
}

func TestInit_MissingPalette(t *testing.T) {
	dir := t.TempDir()

	paths := env.Paths{
		PalettesDir:        dir,
		CurrentPaletteFile: filepath.Join(dir, "current.yaml"),
	}

	err := Init("missing", paths)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestInit_Overwrite(t *testing.T) {
	dir := t.TempDir()

	palettesDir := filepath.Join(dir, "palettes")
	os.MkdirAll(palettesDir, 0755)

	dark := `base00: "000001"
base01: "111111"
base02: "222222"
base03: "333333"
base04: "444444"
base05: "555555"
base06: "666666"
base07: "777777"
base08: "888888"
base09: "999999"
base0A: "aaaaaa"
base0B: "bbbbbb"
base0C: "cccccc"
base0D: "dddddd"
base0E: "eeeeee"
base0F: "ffffff"
`

	writeFile(t, filepath.Join(palettesDir, "dark.yaml"), dark)

	currentPath := filepath.Join(palettesDir, "current.yaml")

	old := `base00: "000000"
base01: "111111"
base02: "222222"
base03: "333333"
base04: "444444"
base05: "555555"
base06: "666666"
base07: "777777"
base08: "888888"
base09: "999999"
base0A: "aaaaaa"
base0B: "bbbbbb"
base0C: "cccccc"
base0D: "dddddd"
base0E: "eeeeee"
base0F: "ffffff"
`

	writeFile(t, currentPath, old)

	paths := env.Paths{
		PalettesDir:        palettesDir,
		CurrentPaletteFile: currentPath,
	}

	err := Init("dark", paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	current := readFile(t, currentPath)

	if current != dark {
		t.Errorf("current palette not overwritten:\n%s", current)
	}
}
