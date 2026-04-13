// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

// Map transforms a palette in a map
func (p Palette) Map() map[string]string {
	return map[string]string{
		"base00": p.Base00,
		"base01": p.Base01,
		"base02": p.Base02,
		"base03": p.Base03,
		"base04": p.Base04,
		"base05": p.Base05,
		"base06": p.Base06,
		"base07": p.Base07,
		"base08": p.Base08,
		"base09": p.Base09,
		"base0A": p.Base0A,
		"base0B": p.Base0B,
		"base0C": p.Base0C,
		"base0D": p.Base0D,
		"base0E": p.Base0E,
		"base0F": p.Base0F,
	}
}
