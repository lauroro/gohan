// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"fmt"
	"path/filepath"

	"codeberg.org/roro/gohan/src/env"
)

// loadPaletteFromContent applies the full parsing pipeline to raw YAML content.
//
// It performs:
// - parsing
// - normalization
// - validation
//
// Returns a fully validated Palette or an error.
func loadPaletteFromContent(content string) (Palette, error) {
	p, err := parsePalette(content)
	if err != nil {
		return Palette{}, err
	}

	p = normalizePalette(p)

	if err := validatePalette(p); err != nil {
		return Palette{}, err
	}

	return p, nil
}

// LoadPalette loads a palette by name from the palettes directory.
//
// It reads the YAML file, then applies parsing, normalization,
// and validation.
//
// Returns:
// - parsed and validated Palette
// - raw file content
// - error (if any)
func LoadPalette(name string, paths env.Paths) (Palette, string, error) {
	fileName := name + ".yaml"
	path := filepath.Join(paths.PalettesDir, fileName)

	content, err := readPaletteFile(path)
	if err != nil {
		return Palette{}, "", fmt.Errorf("could not load palette %s: %w", name, err)
	}

	p, err := loadPaletteFromContent(content)
	if err != nil {
		return Palette{}, "", fmt.Errorf("could not load palette %s: %w", name, err)
	}

	return p, content, nil
}

// LoadCurrentPalette loads the currently active palette.
//
// It reads the palette from paths.CurrentPaletteFile and applies
// the same pipeline used by LoadPalette.
//
// Returns:
// - parsed and validated Palette
// - raw file content
// - error (if any)
func LoadCurrentPalette(paths env.Paths) (Palette, string, error) {
	content, err := readPaletteFile(paths.CurrentPaletteFile)
	if err != nil {
		return Palette{}, "", fmt.Errorf("could not load current palette: %w", err)
	}

	p, err := loadPaletteFromContent(content)
	if err != nil {
		return Palette{}, "", fmt.Errorf("could not load current palette: %w", err)
	}

	return p, content, nil
}
