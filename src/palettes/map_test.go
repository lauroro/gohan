// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import "testing"

func TestPaletteMap_AllKeysPresent(t *testing.T) {
	p := validPalette()

	m := p.Map()

	expectedKeys := []string{
		"base00", "base01", "base02", "base03",
		"base04", "base05", "base06", "base07",
		"base08", "base09", "base0A", "base0B",
		"base0C", "base0D", "base0E", "base0F",
	}

	if len(m) != 16 {
		t.Fatalf("expected 16 entries, got %d", len(m))
	}

	for _, k := range expectedKeys {
		if _, ok := m[k]; !ok {
			t.Errorf("missing key %s", k)
		}
	}
}

func TestPaletteMap_CorrectValues(t *testing.T) {
	p := validPalette()

	m := p.Map()

	tests := map[string]string{
		"base00": "000000",
		"base01": "111111",
		"base0A": "aaaaaa",
		"base0F": "ffffff",
	}

	for k, expected := range tests {
		if got := m[k]; got != expected {
			t.Errorf("expected %s for %s, got %s", expected, k, got)
		}
	}
}

func TestPaletteMap_ConsistencyWithStruct(t *testing.T) {
	p := validPalette()

	m := p.Map()

	if m["base00"] != p.Base00 {
		t.Errorf("base00 mismatch")
	}
	if m["base0F"] != p.Base0F {
		t.Errorf("base0F mismatch")
	}
}
