// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import (
	"fmt"
	"regexp"

	"gohan/src/palettes"
)

var hexRegex = regexp.MustCompile(`(?i)[0-9a-f]{6}`)

// transformBlocks applies a palette transformation to a set of blocks.
//
// It scans each block for hexadecimal color values (6-digit hex strings),
// validates that each found color exists in the "old" palette, and replaces
// it with the corresponding color from the "new" palette.
//
// The transformation works as follows:
//   - Each detected hex value is normalized (lowercased, without '#').
//   - The value is looked up in the reverse map of the old palette (hex -> baseXX).
//   - The corresponding baseXX is then used to fetch the new hex value
//     from the new palette.
//   - Only the matched 6-digit portion is replaiced, preserving surrounding
//     prefixes/suffixes (e.g. "#", "0x", or trailing alpha like "ff").
//
// The function performs two passes:
//  1. Validation: ensures all matched hex values exist in the old palette.
//  2. Transformation: applies replacements.
//
// If any hex value is not found in the old palette, the function returns an error
// and no transformation is applied.
//
// Matches that are part of a larger hex sequence (e.g. embedded in longer strings)
// are skipped unless they represent a valid standalone or prefix hex (e.g. "444444ff").
//
// Returns a new slice of blocks with transformed content, or an error.
func transformBlocks(blocks []block, old palettes.Palette, new palettes.Palette) ([]block, error) {
	oldMap := old.ReverseMap()
	newMap := new.Map()

	// VALIDATION
	for _, b := range blocks {
		matches := hexRegex.FindAllStringIndex(b.Content, -1)

		for _, m := range matches {
			start, end := m[0], m[1]

			if !isValidHexMatch(b.Content, start, end) {
				continue
			}

			hex := b.Content[start:end]
			normalized := normalizeHex(hex)

			if _, ok := oldMap[normalized]; !ok {
				return nil, fmt.Errorf("unknown hex %q in template", hex)
			}
		}
	}

	// TRANSFORMATION
	out := make([]block, len(blocks))

	for i, b := range blocks {
		out[i] = b

		matches := hexRegex.FindAllStringIndex(b.Content, -1)
		result := []byte{}
		last := 0

		for _, m := range matches {
			start, end := m[0], m[1]

			if !isValidHexMatch(b.Content, start, end) {
				continue
			}

			// append previous untuoched content
			result = append(result, b.Content[last:start]...)

			hex := b.Content[start:end]
			normalized := normalizeHex(hex)

			base := oldMap[normalized]
			replacement := newMap[base]

			result = append(result, replacement...)
			last = end
		}

		// append remaining
		result = append(result, b.Content[last:]...)

		out[i].Content = string(result)
	}

	return out, nil
}

func isValidHexMatch(s string, start, end int) bool {
	if start > 0 && isHexChar(s[start-1]) {
		return false
	}
	if end < len(s) && isHexChar(s[end]) {
		if start > 0 && isHexChar(s[start-1]) {
			return false
		}
	}
	return true
}

func isHexChar(c byte) bool {
	return (c >= '0' && c <= '9') ||
		(c >= 'a' && c <= 'f') ||
		(c >= 'A' && c <= 'F')
}

func normalizeHex(h string) string {
	if len(h) > 0 && h[0] == '#' {
		h = h[1:]
	}
	return lower(h)
}

func lower(s string) string {
	out := make([]byte, len(s))
	for i := range s {
		c := s[i]
		if c >= 'A' && c <= 'F' {
			c += 32
		}
		out[i] = c
	}
	return string(out)
}
