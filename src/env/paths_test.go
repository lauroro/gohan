// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package env

import (
	"os"
	"testing"
)

func TestBuildPaths_WithAllEnvVars(t *testing.T) {
	t.Setenv("HOME", "/home/test")
	t.Setenv("XDG_CONFIG_HOME", "/custom/config")
	t.Setenv("XDG_STATE_HOME", "/custom/state")

	result := BuildPaths()

	if result.ConfigFile != "/custom/config/gohan/config.yaml" {
		t.Errorf("unexpected ConfigFile: %s", result.ConfigFile)
	}

	if result.CurrentPaletteFile != "/custom/state/gohan/current.yaml" {
		t.Errorf("unexpected CurrentPaletteFile: %s", result.CurrentPaletteFile)
	}

	if result.PalettesDir != "/custom/config/gohan/palettes" {
		t.Errorf("unexpected PalettesDir: %s", result.PalettesDir)
	}
}

func TestBuildPaths_WithFallbacks(t *testing.T) {
	t.Setenv("HOME", "/home/test")

	os.Unsetenv("XDG_CONFIG_HOME")
	os.Unsetenv("XDG_STATE_HOME")

	result := BuildPaths()

	if result.ConfigFile != "/home/test/.config/gohan/config.yaml" {
		t.Errorf("unexpected ConfigFile: %s", result.ConfigFile)
	}

	if result.CurrentPaletteFile != "/home/test/.local/state/gohan/current.yaml" {
		t.Errorf("unexpected CurrentPaletteFile: %s", result.CurrentPaletteFile)
	}

	if result.PalettesDir != "/home/test/.config/gohan/palettes" {
		t.Errorf("unexpected PalettesDir: %s", result.PalettesDir)
	}
}
func TestBuildPaths_NoCaching(t *testing.T) {
	t.Setenv("HOME", "/home/test")
	t.Setenv("XDG_CONFIG_HOME", "/first")

	first := BuildPaths()

	t.Setenv("XDG_CONFIG_HOME", "/second")

	second := BuildPaths()

	if first == second {
		t.Errorf("expected different results, got same")
	}
}
