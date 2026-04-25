// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package actions

import (
	"gohan/src/env"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestApply_EndToEnd(t *testing.T) {
	dir := t.TempDir()

	palettesDir := filepath.Join(dir, "palettes")
	os.MkdirAll(palettesDir, 0755)

	targetFile := filepath.Join(dir, "file.txt")
	configFile := filepath.Join(dir, "config.yaml")

	// current palette
	writeFile(t, filepath.Join(palettesDir, "current.yaml"), `
base00: "000000"
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
`)

	// new palette
	writeFile(t, filepath.Join(palettesDir, "dark.yaml"), `
base00: "123456"
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
`)

	// target file with template block
	writeFile(t, targetFile, `before
@gohan:start
color: 000000
@gohan:end
after`)

	// config
	writeFile(t, configFile, `
targets:
  - `+targetFile+`
`)

	paths := env.Paths{
		Home:               dir,
		PalettesDir:        palettesDir,
		ConfigFile:         configFile,
		CurrentPaletteFile: filepath.Join(palettesDir, "current.yaml"),
	}

	err := Apply("dark", paths)
	if err != nil {
		t.Fatalf("apply failed: %v", err)
	}

	out := readFile(t, targetFile)

	current := readFile(t, filepath.Join(palettesDir, "current.yaml"))

	if !strings.Contains(current, "123456") {
		t.Errorf("current palette was not updated:\n%s", current)
	}

	if !strings.Contains(out, "123456") {
		t.Errorf("palette was not applied correctly:\n%s", out)
	}

	if !strings.Contains(out, "@gohan:start") || !strings.Contains(out, "@gohan:end") {
		t.Errorf("markers must be preserved:\n%s", out)
	}

	if !strings.Contains(out, "before") || !strings.Contains(out, "after") {
		t.Errorf("outside content must be preserved:\n%s", out)
	}
}
