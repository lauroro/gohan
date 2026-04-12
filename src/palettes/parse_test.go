// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import "testing"

func TestParsePalette_Full(t *testing.T) {
	input := `
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

	p, err := parsePalette(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p.Base00 != "#000000" || p.Base0F != "#ffffff" {
		t.Errorf("unexpected values: %+v", p)
	}
}

func TestParsePalette_WithoutHash(t *testing.T) {
	input := `base00: abcdef`

	p, err := parsePalette(input)
	if err != nil {
		t.Fatal(err)
	}

	if p.Base00 != "abcdef" {
		t.Errorf("expected abcdef, got %s", p.Base00)
	}
}

func TestParsePalette_WithHash(t *testing.T) {
	input := `base00: "#abcdef"`

	p, err := parsePalette(input)
	if err != nil {
		t.Fatal(err)
	}

	if p.Base00 != "#abcdef" {
		t.Errorf("expected #abcdef, got %s", p.Base00)
	}
}

func TestParsePalette_Quoted(t *testing.T) {
	input := `base00: "abcdef"`

	p, err := parsePalette(input)
	if err != nil {
		t.Fatal(err)
	}

	if p.Base00 != "abcdef" {
		t.Errorf("expected abcdef, got %s", p.Base00)
	}
}

func TestParsePalette_InvalidYAML(t *testing.T) {
	input := `::: invalid yaml :::`

	_, err := parsePalette(input)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
