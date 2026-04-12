// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"os"
	"testing"
)

func TestReadPaletteFile_Success(t *testing.T) {
	tmp := t.TempDir()
	file := tmp + "/palette.txt"

	err := os.WriteFile(file, []byte("hello"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	result, err := ReadPaletteFile(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != "hello" {
		t.Errorf("expected hello, got %s", result)
	}
}

func TestReadPaletteFile_NotFound(t *testing.T) {
	_, err := ReadPaletteFile("/non/existent/file")
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
