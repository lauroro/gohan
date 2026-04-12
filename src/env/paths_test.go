// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package env

import (
	"os"
	"sync"
	"testing"
)

func resetPaths() {
	paths = Paths{}
	once = sync.Once{}
}

func TestGetPaths_WithAllEnvVars(t *testing.T) {
	resetPaths()

	os.Setenv("HOME", "/home/test")
	os.Setenv("XDG_CONFIG_HOME", "/custom/config")
	os.Setenv("XDG_STATE_HOME", "/custom/state")
	defer os.Clearenv()

	result := GetPaths()

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

func TestGetPaths_WithFallbacks(t *testing.T) {
	resetPaths()

	os.Setenv("HOME", "/home/test")
	os.Unsetenv("XDG_CONFIG_HOME")
	os.Unsetenv("XDG_STATE_HOME")
	defer os.Clearenv()

	result := GetPaths()

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

func TestGetPaths_OnceBehavior(t *testing.T) {
	resetPaths()

	os.Setenv("HOME", "/home/test")
	os.Setenv("XDG_CONFIG_HOME", "/first")
	defer os.Clearenv()

	first := GetPaths()

	os.Setenv("XDG_CONFIG_HOME", "/second")

	second := GetPaths()

	if first != second {
		t.Errorf("expected same result due to sync.Once")
	}
}
