// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"fmt"
	"path/filepath"

	"gohan/src/env"
	"gohan/src/tools"
)

// readPalette reads a palette file by name from the configured palettes directory.
// It returns the file content as a string or an error if the file cannot be read.
func readPalette(name string, paths env.Paths) (string, error) {
	path := filepath.Join(paths.PalettesDir, name)

	content, err := tools.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not read palette %s: %w", name, err)
	}

	return content, nil
}

// readCurrentPalette reads the currently active palette file from the configured path.
// It returns the file content as a string or an error if the file cannot be read.
func readCurrentPalette(paths env.Paths) (string, error) {
	content, err := tools.ReadFile(paths.CurrentPaletteFile)
	if err != nil {
		return "", fmt.Errorf("could not read current palette: %w", err)
	}

	return content, nil
}
