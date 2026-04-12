// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package env

type Paths struct {
	ConfigFile         string
	CurrentPaletteFile string
	PalettesDir        string
}

// BuildPaths constructs application-specific filesystem paths based on
// environment variables following the XDG Base Directory specification.
//
// It reads:
// - $HOME (required, assumed to be non-empty)
// - $XDG_CONFIG_HOME (defaults to $HOME/.config if unset or empty)
// - $XDG_STATE_HOME (defaults to $HOME/.local/state if unset or empty)
//
// The returned Paths struct contains:
// - ConfigFile: path to the main configuration file
// - CurrentPaletteFile: path to the currently active palette file
// - PalettesDir: directory containing available palettes
//
// This funtion is pure and does not cache results, making it safe and
// predictable for use in tests and dependency injection scenarios.
func BuildPaths() Paths {
	home := env("HOME", "")
	xdgConfigDir := env("XDG_CONFIG_HOME", home+"/.config")
	xdgStateDir := env("XDG_STATE_HOME", home+"/.local/state")

	gohanDir := xdgConfigDir + "/gohan"

	return Paths{
		ConfigFile:         gohanDir + "/config.yaml",
		CurrentPaletteFile: xdgStateDir + "/gohan/current.yaml",
		PalettesDir:        gohanDir + "/palettes",
	}
}
