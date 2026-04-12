// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

// validatePalette checks that a palette is complete and valid.
// It ensures:
// - all 16 base colors are present
// - all values are valid 6-char hex strings (lowercase, no '#')
func validatePalette(p Palette) error
