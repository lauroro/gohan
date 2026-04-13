// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"fmt"
	"regexp"
)

var hexRegex = regexp.MustCompile(`^[0-9a-f]{6}$`)

// validatePalette checks that a palette is complete and valid.
// It ensures:
// - all 16 base colors are present
// - all values are valid 6-char hex strings (lowercase, no '#')
// - all values are unique (no duplicates)
func validatePalette(p Palette) error {
	values := map[string]string{
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

	seen := make(map[string]string)

	for base, hex := range values {
		if hex == "" {
			return fmt.Errorf("invalid palette: %s is empty", base)
		}

		if !hexRegex.MatchString(hex) {
			return fmt.Errorf("invalid palette: %s has invalid hex value %q", base, hex)
		}

		if prev, exists := seen[hex]; exists {
			return fmt.Errorf(
				"invalid palette: duplicate hex %q found in %s and %s",
				hex, prev, base,
			)
		}

		seen[hex] = base
	}

	return nil
}
