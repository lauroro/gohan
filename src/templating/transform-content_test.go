// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import (
	"strings"
	"testing"

	"codeberg.org/roro/gohan/src/palettes"
)

var old = palettes.Palette{
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

var new = palettes.Palette{
	Base00: "1d2021",
	Base01: "3c3836",
	Base02: "504945",
	Base03: "665c54",
	Base04: "bdae93",
	Base05: "d5c4a1",
	Base06: "ebdbb2",
	Base07: "fbf1c7",
	Base08: "fb4934",
	Base09: "fe8019",
	Base0A: "fabd2f",
	Base0B: "b8bb26",
	Base0C: "8ec07c",
	Base0D: "83a598",
	Base0E: "d3869b",
	Base0F: "d65d0e",
}

func TestTransformContent_SingleBlock(t *testing.T) {
	input := `before
@gohan:start
color: 000000
@gohan:end
after`

	out, err := transformContent(input, old, new)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(out, "1d2021") {
		t.Fatalf("expected transformed content, got:\n%s", out)
	}
}

func TestTransformContent_PreserveFormatting(t *testing.T) {
	input := `before
# @gohan:start
color: 000000
# @gohan:end
after`

	expected := `before
# @gohan:start
color: 1d2021
# @gohan:end
after`

	out, err := transformContent(input, old, new)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out != expected {
		t.Fatalf("output format was not as expected:\n%s", out)
	}
}

func TestTransformContent_ComplexBlock(t *testing.T) {
	input := `before
// @gohan:start
with_zero: 0x000000
with_hash: #111111
with_alpha = eeeeeeff
this_one_is_hard = 000000, 111111, rgb(222222)
-- @gohan:end
after`

	out, err := transformContent(input, old, new)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(out, "with_zero: 0x1d2021") || !strings.Contains(out, "with_hash: #3c3836") || !strings.Contains(out, "with_alpha = d3869bff") || !strings.Contains(out, "this_one_is_hard = 1d2021, 3c3836, rgb(504945)") {
		t.Fatalf("expected transformed content, got:\n%s", out)
	}
}

func TestTransformContent_MultipleBlocks(t *testing.T) {
	input := `@gohan:start
color: 000000
@gohan:end

@gohan:start
color: 111111
@gohan:end`

	out, err := transformContent(input, old, new)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(out, "1d2021") || !strings.Contains(out, "3c3836") {
		t.Fatalf("blocks not transformed correctly:\n%s", out)
	}
}

func TestTransformContent_PreservesComments(t *testing.T) {
	input := `/* @gohan:start */
color: 000000
/* @gohan:end */

# @gohan:start
color: 000000
# @gohan:end

// @gohan:start
color: 000000
// @gohan:end`

	out, err := transformContent(input, old, new)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(out, "/* @gohan:start */") ||
		!strings.Contains(out, "# @gohan:start") ||
		!strings.Contains(out, "// @gohan:start") {
		t.Fatalf("markers not preserved:\n%s", out)
	}

	if !strings.Contains(out, "1d2021") {
		t.Fatalf("expected transformed content")
	}
}

func TestTransformContent_NoBlocks(t *testing.T) {
	input := `no blocks here`

	out, err := transformContent(input, testPalette(), testPalette())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out != input {
		t.Fatalf("expected unchanged content")
	}
}

func TestTransformContent_InvalidHex(t *testing.T) {
	input := `@gohan:start
color: deadbe
@gohan:end`

	_, err := transformContent(input, testPalette(), testPalette())
	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestTransformContent_PreservesOutsideText(t *testing.T) {
	input := `AAA
@gohan:start
color: 000000
@gohan:end
BBB`

	out, err := transformContent(input, old, new)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(out, "AAA") || !strings.Contains(out, "BBB") {
		t.Fatalf("outside text not preserved:\n%s", out)
	}
}
