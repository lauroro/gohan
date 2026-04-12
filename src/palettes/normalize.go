// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import "strings"

// normalizePalette normalizes all palette color values by:
// - removing the leading '#' if present
// - convetring all characters to lowercase
func normalizePalette(p Palette) Palette {
	p.Base00 = normalizeHex(p.Base00)
	p.Base01 = normalizeHex(p.Base01)
	p.Base02 = normalizeHex(p.Base02)
	p.Base03 = normalizeHex(p.Base03)
	p.Base04 = normalizeHex(p.Base04)
	p.Base05 = normalizeHex(p.Base05)
	p.Base06 = normalizeHex(p.Base06)
	p.Base07 = normalizeHex(p.Base07)
	p.Base08 = normalizeHex(p.Base08)
	p.Base09 = normalizeHex(p.Base09)
	p.Base0A = normalizeHex(p.Base0A)
	p.Base0B = normalizeHex(p.Base0B)
	p.Base0C = normalizeHex(p.Base0C)
	p.Base0D = normalizeHex(p.Base0D)
	p.Base0E = normalizeHex(p.Base0E)
	p.Base0F = normalizeHex(p.Base0F)

	return p
}

func normalizeHex(s string) string {
	s = strings.TrimPrefix(s, "#")
	return strings.ToLower(s)
}
