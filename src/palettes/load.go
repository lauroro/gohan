// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"fmt"
	"path/filepath"

	"gohan/src/env"
)

// LoadPalette loads a palette by name from the palettes directory.
//
// The function performs the full pipeline:
// - reads the YAML file from disk
// - parses it into a Palette struct
// - normalizes hex values
// - validates result
//
// It returns a fully validated Palette or an error if any step fails.
func LoadPalette(name string, paths env.Paths) (Palette, error) {
	fileName := name + ".yaml"
	path := filepath.Join(paths.PalettesDir, fileName)

	content, err := readPaletteFile(path)
	if err != nil {
		return Palette{}, fmt.Errorf("could not load palette %s: %w", name, err)
	}

	p, err := parsePalette(content)
	if err != nil {
		return Palette{}, fmt.Errorf("could not load palette %s: %w", name, err)
	}

	p = normalizePalette(p)

	if err := validatePalette(p); err != nil {
		return Palette{}, fmt.Errorf("could not load palette %s: %w", name, err)
	}

	return p, nil
}
