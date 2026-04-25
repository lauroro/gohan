// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package actions

import (
	"os"
	"path/filepath"
	"testing"

	"gohan/src/env"
)

func writePalette(t *testing.T, dir, name string) {
	t.Helper()
	_ = os.WriteFile(filepath.Join(dir, name+".yaml"), []byte(""), 0644)
}

func TestList_Empty(t *testing.T) {
	dir := t.TempDir()

	paths := env.Paths{
		PalettesDir: dir,
	}

	list, err := List(paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(list) != 0 {
		t.Errorf("expected empty list, got %v", list)
	}
}

func TestList_WithPalettes(t *testing.T) {
	dir := t.TempDir()

	writePalette(t, dir, "dracula")
	writePalette(t, dir, "nord")

	paths := env.Paths{
		PalettesDir: dir,
	}

	list, err := List(paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := map[string]bool{
		"dracula": true,
		"nord":    true,
	}

	if len(list) != 2 {
		t.Fatalf("expected 2 palettes, got %v", list)
	}

	for _, p := range list {
		if !expected[p] {
			t.Errorf("unexpected palette: %s", p)
		}
	}
}
