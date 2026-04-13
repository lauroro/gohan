// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriteFile writes the given content to the specified file path.
//
// If the file does not exist it will be created with permission 0644.
// If it already exists it will be truncated and overwritten.
//
// If the parent directories do not exist they will be created automatically.
//
// It returns an error if the file cannot be written.
func WriteFile(path, content string) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("could not create directories for %s: %w", path, err)
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("could not write file %s: %w", path, err)
	}

	return nil
}
