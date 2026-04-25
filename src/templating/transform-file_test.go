// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import (
	"os"
	"path/filepath"
	"testing"

	"gohan/src/palettes"
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

func testPalette() palettes.Palette {
	return palettes.Palette{
		Base00: "000000",
		Base01: "111111",
		Base02: "222222",
		Base03: "333333",
		Base04: "444444",
		Base05: "555555",
		Base06: "666666",
		Base07: "777777",
		Base08: "888888",
		Base09: "999999",
		Base0A: "aaaaaa",
		Base0B: "bbbbbb",
		Base0C: "cccccc",
		Base0D: "dddddd",
		Base0E: "eeeeee",
		Base0F: "ffffff",
	}
}

func TestTransformFile_Success(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "test.txt")

	input := `before
@gohan:start
color: 000000
@gohan:end
after`

	writeFile(t, file, input)

	old := testPalette()
	new := testPalette()
	new.Base00 = "123456"

	err := TransformFile(file, old, new)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := readFile(t, file)

	if !contains(out, "123456") {
		t.Errorf("expected transformed content, got:\n%s", out)
	}
}

func TestTransformFile_NoBlocks(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "test.txt")

	input := "no blocks here"
	writeFile(t, file, input)

	err := TransformFile(file, testPalette(), testPalette())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := readFile(t, file)

	if out != input {
		t.Errorf("file should remain unchanged")
	}
}

func TestTransformFile_FileMissing(t *testing.T) {
	err := TransformFile("missing.txt", testPalette(), testPalette())

	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestTransformFile_InvalidHex(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "test.txt")

	input := `@gohan:start
color: deadbe
@gohan:end`

	writeFile(t, file, input)

	err := TransformFile(file, testPalette(), testPalette())

	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
