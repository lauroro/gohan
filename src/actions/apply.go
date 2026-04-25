// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package actions

import (
	"fmt"

	"gohan/src/config"
	"gohan/src/env"
	"gohan/src/palettes"
	"gohan/src/templating"

	"gopkg.in/yaml.v3"
)

// Apply applies a palette to all configured targets.
//
// It performs the following steps:
// - loads the current palette
// - loads the new palette by name
// - loads the configuration file
// - transforms all target files
// - executes configured commands
//
// It returns an error if any step fails.
func Apply(paletteName string, paths env.Paths) error {
	currentPalette, err := palettes.LoadCurrentPalette(paths)
	if err != nil {
		return fmt.Errorf("apply: %w", err)
	}

	newPalette, err := palettes.LoadPalette(paletteName, paths)
	if err != nil {
		return fmt.Errorf("apply: %w", err)
	}

	cfg, err := config.LoadConfig(paths)
	if err != nil {
		return fmt.Errorf("apply: %w", err)
	}

	for _, t := range cfg.Targets {
		if err := templating.TransformFile(t, currentPalette, newPalette); err != nil {
			return fmt.Errorf("apply: %w", err)
		}
	}

	newCurrentPalette, err := yaml.Marshal(newPalette)
	if err != nil {
		return fmt.Errorf("apply: %w", err)
	}

	if err := palettes.SaveCurrentPalette(string(newCurrentPalette), paths); err != nil {
		return fmt.Errorf("apply: %w", err)
	}

	for _, c := range cfg.Commands {
		if _, err := config.ExecCmd(c); err != nil {
			return fmt.Errorf("apply: %w", err)
		}
	}

	return nil
}
