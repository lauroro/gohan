// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package tools

import (
	"fmt"
	"os"
)

// ReadFile reads the entire content of the file at the given path
// and returns it as a string.
//
// It returns an error if the file cannot be read.
func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not read file %s: %w", path, err)
	}

	return string(data), nil
}
