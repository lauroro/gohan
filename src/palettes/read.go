// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"fmt"

	"gohan/src/tools"
)

// ReadPaletteFile reads a palette file from the given path.
// It return the file content as a string or an error if the file cannot be read.
func ReadPaletteFile(path string) (string, error) {
	content, err := tools.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not read palette file %s: %w", path, err)
	}
	return content, nil
}
