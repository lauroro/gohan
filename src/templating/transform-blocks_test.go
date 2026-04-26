// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import (
	"testing"

	"codeberg.org/roro/gohan/src/palettes"
)

func contains(s, substr string) bool {
	return indexOf(s, substr) >= 0
}

func indexOf(s, substr string) int {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func oldPalette() palettes.Palette {
	return palettes.Palette{
		Base00: "000000",
		Base01: "123456",
		Base02: "abcdef",
		Base03: "444444",
		Base04: "555555",
		Base05: "666666",
		Base06: "777777",
		Base07: "888888",
		Base08: "999999",
		Base09: "aaaaaa",
		Base0A: "bbbbbb",
		Base0B: "cccccc",
		Base0C: "dddddd",
		Base0D: "eeeeee",
		Base0E: "ffffff",
		Base0F: "121212",
	}
}

func newPalette() palettes.Palette {
	return palettes.Palette{
		Base00: "000001",
		Base01: "123457",
		Base02: "abcdee",
		Base03: "444445",
		Base04: "555556",
		Base05: "666667",
		Base06: "777778",
		Base07: "888889",
		Base08: "999990",
		Base09: "aaaaab",
		Base0A: "bbbbbc",
		Base0B: "cccccd",
		Base0C: "ddddde",
		Base0D: "eeeee0",
		Base0E: "fffff1",
		Base0F: "121213",
	}
}

func TestTransformBlocks_FullExample(t *testing.T) {
	blocks := []block{
		{
			Content: `
background-color: #abcdef
something = 0x123456
another: 444444ff
pippo = 000000
`,
		},
	}

	out, err := transformBlocks(blocks, oldPalette(), newPalette())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := out[0].Content

	expected := []string{
		"#abcdee",
		"0x123457",
		"444445ff",
		"000001",
	}

	for _, exp := range expected {
		if !contains(got, exp) {
			t.Errorf("expected %q in output, got:\n%s", exp, got)
		}
	}
}

func TestTransformBlocks_SingleHex(t *testing.T) {
	blocks := []block{
		{Content: "color: 000000"},
	}

	out, err := transformBlocks(blocks, oldPalette(), newPalette())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !contains(out[0].Content, "000001") {
		t.Errorf("expected transformation to 000001, got %q", out[0].Content)
	}
}

func TestTransformBlocks_MultipleBlocks(t *testing.T) {
	blocks := []block{
		{Content: "a: 000000"},
		{Content: "b: abcdef"},
	}

	out, err := transformBlocks(blocks, oldPalette(), newPalette())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(out) != 2 {
		t.Fatalf("expected 2 blocks, got %d", len(out))
	}

	if !contains(out[0].Content, "000001") {
		t.Errorf("first block not transformed")
	}

	if !contains(out[1].Content, "abcdee") {
		t.Errorf("second block not transformed")
	}
}

func TestTransformBlocks_CaseInsensitive(t *testing.T) {
	blocks := []block{
		{Content: "color: ABCDEF"},
	}

	out, err := transformBlocks(blocks, oldPalette(), newPalette())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !contains(out[0].Content, "abcdee") {
		t.Errorf("expected lowercase transformed value, got %q", out[0].Content)
	}
}

func TestTransformBlocks_NoHex(t *testing.T) {
	blocks := []block{
		{Content: "hello world"},
	}

	out, err := transformBlocks(blocks, oldPalette(), newPalette())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out[0].Content != "hello world" {
		t.Errorf("content should remain unchanged, got %q", out[0].Content)
	}
}

func TestTransformBlocks_UnknownHex(t *testing.T) {
	blocks := []block{
		{Content: "color: 999998"},
	}

	_, err := transformBlocks(blocks, oldPalette(), newPalette())

	if err == nil {
		t.Fatal("expected error for unknown hex, got nil")
	}
}
