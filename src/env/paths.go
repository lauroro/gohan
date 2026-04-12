// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package env

import "sync"

type Paths struct {
	ConfigFile         string
	CurrentPaletteFile string
	PalettesDir        string
}

var (
	paths Paths
	once  sync.Once
)

// GetPaths returns application-specific filesystem paths derived from
// environment variables following the XDG Base Directory specification.
//
// The paths are computed only once using sync.Once and then cached for
// subseqeuent calls.
//
// It uses:
// - $HOME (required, panics if missing)
// - $XDG_CONFIG_HOME (defaults to $HOME/.config)
// - $XDG_STATE_HOME (defaults to $HOME/.local/state)
//
// The returned Paths struct contains locations for the config file,
// current palette file, and palettes directory.
func GetPaths() Paths {
	once.Do(func() {
		home := env("HOME", "")
		xdgConfigDir := env("XDG_CONFIG_HOME", home+"/.config")
		gohanDir := xdgConfigDir + "/gohan"
		xdgStateDir := env("XDG_STATE_HOME", home+"/.local/state")

		paths = Paths{
			ConfigFile:         gohanDir + "/config.yaml",
			CurrentPaletteFile: xdgStateDir + "/gohan/current.yaml",
			PalettesDir:        gohanDir + "/palettes",
		}
	})
	return paths
}
