// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import (
	"fmt"
	"regexp"
	"strings"

	"codeberg.org/roro/gohan/src/palettes"
)

const (
	startMarker = "@gohan:start"
	endMarker   = "@gohan:end"
)

var hexRegex = regexp.MustCompile(`(?i)[0-9a-f]{6}`)

// transformContent applies a palette transformation to all templated blocks
// found in the provided content string.
//
// A block is any sequence of lines between a line containing "@gohan:start"
// and a line containing "@gohan:end". The markers themselves are preserved
// exactly as they appear (including comments or prefixes).
//
// Behavior:
//   - Lines outside of blocks are copied unchanged.
//   - Lines inside blocks are scanned for 6-digit hexadecimal values.
//   - Each hex value is validated against the "old" palette and replaced
//     with the corresponding value from the "new" palette.
//   - Only valid standalone hex matches are transformed.
//   - If any hex value inside a block is not found in the old palette, an error is returned.
//
// Limitations:
//   - Inline blocks on a single line are not supported.
//   - Nested blocks are not supported.
func transformContent(content string, old, new palettes.Palette) (string, error) {
	lines := strings.Split(content, "\n")

	oldMap := old.ReverseMap()
	newMap := new.Map()

	var result strings.Builder
	inBlock := false

	for i, line := range lines {
		switch {
		case strings.Contains(line, startMarker):
			inBlock = true
			result.WriteString(line)

		case strings.Contains(line, endMarker):
			inBlock = false
			result.WriteString(line)

		case inBlock:
			matches := hexRegex.FindAllStringIndex(line, -1)

			var out []byte
			last := 0

			for _, m := range matches {
				start, end := m[0], m[1]

				if !isValidHexMatch(line, start, end) {
					continue
				}

				out = append(out, line[last:start]...)

				hex := line[start:end]
				normalized := normalizeHex(hex)

				base, ok := oldMap[normalized]
				if !ok {
					return "", fmt.Errorf("unknown hex %q in template", hex)
				}

				replacement := newMap[base]
				out = append(out, replacement...)

				last = end
			}

			out = append(out, line[last:]...)
			result.WriteString(string(out))

		default:
			result.WriteString(line)
		}

		if i < len(lines)-1 {
			result.WriteByte('\n')
		}
	}

	return result.String(), nil
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
