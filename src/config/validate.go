// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import "fmt"

// validateConfig checks that a Config struct is valid.
//
// Rules:
// - Targets must be a non-empty slice of non-empty strings
// - Commands is optional, but if present must contain only non-empty strings
func validateConfig(c Config) error {
	if len(c.Targets) == 0 {
		return fmt.Errorf("invalid config: targets must not be empty")
	}

	for i, t := range c.Targets {
		if t == "" {
			return fmt.Errorf("invalid config: targets[%d] is empty", i)
		}
	}

	for i, cmd := range c.Commands {
		if cmd == "" {
			return fmt.Errorf("invalid config: commands[%d] is empty", i)
		}
	}

	return nil
}
