// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"codeberg.org/roro/gohan/src/env"
	"fmt"
	"os"
	"strings"
)

// ListPalettes returns the list of available palette names.
//
// It scans the directory defined in paths.PalettesDir and returns
// all files ending with ".yaml", stripping the extension.
//
// Only regular files with ".yaml" suffix are considered.
// The returned names do not include the ".yaml" extension.
//
// Example:
//
//	"gruvbox.yaml" -> "gruvbox"
func ListPalettes(paths env.Paths) ([]string, error) {
	var list []string

	entries, err := os.ReadDir(paths.PalettesDir)
	if err != nil {
		return nil, fmt.Errorf("could not list palettes: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()

		if before, ok := strings.CutSuffix(name, ".yaml"); ok {
			list = append(list, before)
		}
	}

	return list, nil
}
