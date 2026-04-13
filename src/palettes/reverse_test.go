// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import "testing"

func validPalette() Palette {
	return Palette{
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

func TestReverseMap_AllEntries(t *testing.T) {
	p := validPalette()

	m := p.ReverseMap()

	if len(m) != 16 {
		t.Fatalf("expected 16 entries, got %d", len(m))
	}
}

func TestReverseMap_CorrectMapping(t *testing.T) {
	p := validPalette()

	m := p.ReverseMap()

	tests := map[string]string{
		"000000": "base00",
		"111111": "base01",
		"aaaaaa": "base0A",
		"ffffff": "base0F",
	}

	for hex, expected := range tests {
		if got := m[hex]; got != expected {
			t.Errorf("expected %s for %s, got %s", expected, hex, got)
		}
	}
}

func TestReverseMap_AllBasesPresent(t *testing.T) {
	p := validPalette()

	m := p.ReverseMap()

	expected := map[string]bool{
		"base00": false,
		"base01": false,
		"base02": false,
		"base03": false,
		"base04": false,
		"base05": false,
		"base06": false,
		"base07": false,
		"base08": false,
		"base09": false,
		"base0A": false,
		"base0B": false,
		"base0C": false,
		"base0D": false,
		"base0E": false,
		"base0F": false,
	}

	for _, base := range m {
		expected[base] = true
	}

	for base, found := range expected {
		if !found {
			t.Errorf("missing base %s in reverse map", base)
		}
	}
}
