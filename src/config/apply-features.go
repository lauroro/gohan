// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import "fmt"

// ApplyFeatures applies the enabled features.
// Missing features are the same as disabled features.
func ApplyFeatures(c Config, colorMode string) error {
	if c.Features.GTKAutoColorScheme {
		var scheme string

		switch colorMode {
		case "dark":
			scheme = "prefer-dark"
		case "light":
			scheme = "prefer-light"
		default:
			return fmt.Errorf("unsupported color mode %q", colorMode)
		}

		command := fmt.Sprintf(
			`gsettings set org.gnome.desktop.interface color-scheme %q`,
			scheme,
		)

		if _, err := ExecCmd(command); err != nil {
			return fmt.Errorf("could not set gtk color scheme: %w", err)
		}
	}

	return nil
}
