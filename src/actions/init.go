// SPDX-FileCopyrightText: 2026 lauroro
// SPDX-License-Identifier: GPL-3.0-or-later

package actions

import (
	"fmt"
	"github.com/lauroro/gohan/src/env"
	"github.com/lauroro/gohan/src/palettes"
)

// Init initializes the current palette.
//
// It loads a palette by name and sets it as the current palette
// by copying its raw content into the current palette file.
//
// This is typically used during first-time setup.
func Init(paletteName string, paths env.Paths) error {
	_, newPaletteRaw, err := palettes.LoadPalette(paletteName, paths)
	if err != nil {
		return fmt.Errorf("apply: %w", err)
	}

	if err := palettes.SaveCurrentPalette(newPaletteRaw, paths); err != nil {
		return fmt.Errorf("apply: %w", err)
	}

	return nil
}
