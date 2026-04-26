// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"os"
	"path/filepath"
	"testing"

	"codeberg.org/roro/gohan/src/env"
)

func TestListPalettes_Basic(t *testing.T) {
	dir := t.TempDir()

	files := []string{
		"one.yaml",
		"two.yaml",
		"ignore.txt",
		"three.yaml.bak",
	}

	for _, f := range files {
		err := os.WriteFile(filepath.Join(dir, f), []byte("x"), 0644)
		if err != nil {
			t.Fatalf("failed to create file: %v", err)
		}
	}

	paths := env.Paths{
		PalettesDir: dir,
	}

	list, err := ListPalettes(paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := map[string]bool{
		"one": false,
		"two": false,
	}

	if len(list) != 2 {
		t.Fatalf("expected 2 palettes, got %d (%v)", len(list), list)
	}

	for _, name := range list {
		if _, ok := expected[name]; !ok {
			t.Errorf("unexpected palette %q", name)
		}
		expected[name] = true
	}

	for name, found := range expected {
		if !found {
			t.Errorf("missing palette %q", name)
		}
	}
}

func TestListPalettes_EmptyDir(t *testing.T) {
	dir := t.TempDir()

	paths := env.Paths{
		PalettesDir: dir,
	}

	list, err := ListPalettes(paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(list) != 0 {
		t.Fatalf("expected empty list, got %v", list)
	}
}

func TestListPalettes_IgnoresDirectories(t *testing.T) {
	dir := t.TempDir()

	err := os.Mkdir(filepath.Join(dir, "subdir.yaml"), 0755)
	if err != nil {
		t.Fatalf("failed to create dir: %v", err)
	}

	err = os.WriteFile(filepath.Join(dir, "valid.yaml"), []byte("x"), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	paths := env.Paths{
		PalettesDir: dir,
	}

	list, err := ListPalettes(paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(list) != 1 || list[0] != "valid" {
		t.Errorf("expected [valid], got %v", list)
	}
}

func TestListPalettes_InvalidDir(t *testing.T) {
	paths := env.Paths{
		PalettesDir: "/non/existent/path",
	}

	_, err := ListPalettes(paths)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
