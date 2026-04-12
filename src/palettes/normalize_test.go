// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import "testing"

func TestNormalizePalette(t *testing.T) {
	input := Palette{
		Base00: "#ABCDEF",
		Base01: "123456",
		Base02: "#abcdef",
		Base03: "ABC123",
	}

	result := normalizePalette(input)

	if result.Base00 != "abcdef" {
		t.Errorf("expected abcdef, got %s", result.Base00)
	}

	if result.Base01 != "123456" {
		t.Errorf("expected 123456, got %s", result.Base01)
	}

	if result.Base02 != "abcdef" {
		t.Errorf("expected abcdef, got %s", result.Base02)
	}

	if result.Base03 != "abc123" {
		t.Errorf("expected abc123, got %s", result.Base03)
	}
}

func TestNormalizePalette_Idempotent(t *testing.T) {
	input := Palette{
		Base00: "abcdef",
	}

	result := normalizePalette(input)

	if result.Base00 != "abcdef" {
		t.Errorf("expected abcdef, got %s", result.Base00)
	}
}

func TestNormalizePalette_Empty(t *testing.T) {
	input := Palette{
		Base00: "",
	}

	result := normalizePalette(input)

	if result.Base00 != "" {
		t.Errorf("expected empty string, got %s", result.Base00)
	}
}
