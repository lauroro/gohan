// SPDX-FileCopyrightText: 2026 lauroro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import (
	"fmt"

	"github.com/lauroro/gohan/src/tools"
)

// readConfigFile reads gohan config file.
// It return the file content as a string or an error if the file cannot be read.
func readConfigFile(path string) (string, error) {
	content, err := tools.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not read config file %s: %w", path, err)
	}
	return content, nil
}
