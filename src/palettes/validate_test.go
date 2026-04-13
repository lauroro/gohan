// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import "testing"

func TestValidatePalette_Valid(t *testing.T) {
	p := Palette{
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

	if err := validatePalette(p); err != nil {
		t.Errorf("expected valid palette, got error: %v", err)
	}
}

func TestValidatePalette_MissingField(t *testing.T) {
	p := Palette{
		Base00: "000000",
	}

	err := validatePalette(p)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestValidatePalette_InvalidHex(t *testing.T) {
	p := Palette{
		Base00: "zzzzzz",
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

	err := validatePalette(p)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestValidatePalette_DuplicateHex(t *testing.T) {
	p := Palette{
		Base00: "abcdef",
		Base01: "abcdef",
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

	err := validatePalette(p)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestValidatePalette_EmptyValue(t *testing.T) {
	p := Palette{
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
		Base0F: "",
	}

	err := validatePalette(p)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
