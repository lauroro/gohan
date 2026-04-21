// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import (
	"gohan/src/env"
	"strings"
)

// normalizeConfig normalizes a Config by resolving user-specific paths.
//
// Currently it replaces a leading "~" in each target with the user's home
// directory provided in paths.Home.
//
// Only the first occurrence of "~" is replaced, assuming it represents the
// home directory shortcut.
//
// Commands are left untouched.
//
// Example:
//
//	"~/.config/file" -> "/home/user/.config/file"
func normalizeConfig(c Config, paths env.Paths) Config {
	for i, t := range c.Targets {
		c.Targets[i] = strings.Replace(t, "~", paths.Home, 1)
	}
	return c
}
