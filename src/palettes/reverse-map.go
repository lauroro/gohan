// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

// ReverseMap returns a map from hex value to base name (e.g. "base00").
func (p Palette) ReverseMap() map[string]string {
	return map[string]string{
		p.Base00: "base00",
		p.Base01: "base01",
		p.Base02: "base02",
		p.Base03: "base03",
		p.Base04: "base04",
		p.Base05: "base05",
		p.Base06: "base06",
		p.Base07: "base07",
		p.Base08: "base08",
		p.Base09: "base09",
		p.Base0A: "base0A",
		p.Base0B: "base0B",
		p.Base0C: "base0C",
		p.Base0D: "base0D",
		p.Base0E: "base0E",
		p.Base0F: "base0F",
	}
}
