// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// parsePalette parses a YAML palette string into a Palette struct.
func parsePalette(content string) (Palette, error) {
	var p Palette

	if err := yaml.Unmarshal([]byte(content), &p); err != nil {
		return Palette{}, fmt.Errorf("could not parse palette: %w", err)
	}

	return p, nil
}
