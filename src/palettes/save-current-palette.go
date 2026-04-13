// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"fmt"
	"gohan/src/env"
	"gohan/src/tools"
)

// SaveCurrentPalette writes the given palette content to the current palette file.
//
// The content is written to the path defined in paths.CurrentPaletteFile.
// If the file does not exist, it will be created.
// If it exists, it will be overwritten.
//
// Parent directoires are created automatically if they do not exist.
//
// It returns an error if the file cannot be written.
func SaveCurrentPalette(content string, paths env.Paths) error {
	if err := tools.WriteFile(paths.CurrentPaletteFile, content); err != nil {
		return fmt.Errorf("could not save current palette: %w", err)
	}
	return nil
}
