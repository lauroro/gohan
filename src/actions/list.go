// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package actions

import (
	"codeberg.org/roro/gohan/src/env"
	"codeberg.org/roro/gohan/src/palettes"
	"fmt"
)

// List returns all available palettes from the palettes directory.
//
// It retrieves the list of palette names from the filesystem and
// prints each one to stdout.
//
// Returns an error if palette discovery fails.
func List(paths env.Paths) ([]string, error) {
	list, err := palettes.ListPalettes(paths)
	if err != nil {
		return []string{}, fmt.Errorf("list: %w", err)
	}

	return list, nil
}
