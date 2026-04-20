// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// parseConfig parses a YAML config string into a Config struct
func parseConfig(content string) (Config, error) {
	var c Config

	if err := yaml.Unmarshal([]byte(content), &c); err != nil {
		return Config{}, fmt.Errorf("count not parse config: %w", err)
	}

	return c, nil
}
