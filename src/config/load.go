// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import (
	"codeberg.org/roro/gohan/src/env"
	"fmt"
)

// LoadConfig loads and prepares the application configuration.
//
// It performs the full configuration pipeline:
// - reads the config file from disk
// - parses the YAML content into a Config struct
// - normalizes values (e.g. path expansion)
// - validates the final structure
//
// It returns a valid Config or an error if any step fails.
func LoadConfig(paths env.Paths) (Config, error) {
	content, err := readConfigFile(paths.ConfigFile)
	if err != nil {
		return Config{}, fmt.Errorf("could not load config: %w", err)
	}

	c, err := parseConfig(content)
	if err != nil {
		return Config{}, err
	}

	c = normalizeConfig(c, paths)

	if err := validateConfig(c); err != nil {
		return Config{}, err
	}

	return c, nil
}
