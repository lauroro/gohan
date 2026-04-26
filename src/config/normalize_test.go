// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import (
	"testing"

	"codeberg.org/roro/gohan/src/env"
)

func TestNormalizeConfig_ReplacesHome(t *testing.T) {
	c := Config{
		Targets: []string{
			"~/file",
			"~/dir/file",
		},
	}

	paths := env.Paths{
		Home: "/home/test",
	}

	out := normalizeConfig(c, paths)

	expected := []string{
		"/home/test/file",
		"/home/test/dir/file",
	}

	for i := range expected {
		if out.Targets[i] != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], out.Targets[i])
		}
	}
}

func TestNormalizeConfig_NoTilde(t *testing.T) {
	c := Config{
		Targets: []string{
			"/etc/config",
			"/var/log",
		},
	}

	paths := env.Paths{
		Home: "/home/test",
	}

	out := normalizeConfig(c, paths)

	for i := range c.Targets {
		if out.Targets[i] != c.Targets[i] {
			t.Errorf("expected %q, got %q", c.Targets[i], out.Targets[i])
		}
	}
}

func TestNormalizeConfig_OnlyFirstTilde(t *testing.T) {
	c := Config{
		Targets: []string{
			"~/dir/~/file",
		},
	}

	paths := env.Paths{
		Home: "/home/test",
	}

	out := normalizeConfig(c, paths)

	expected := "/home/test/dir/~/file"

	if out.Targets[0] != expected {
		t.Errorf("expected %q, got %q", expected, out.Targets[0])
	}
}

func TestNormalizeConfig_EmptyTargets(t *testing.T) {
	c := Config{
		Targets: []string{},
	}

	paths := env.Paths{
		Home: "/home/test",
	}

	out := normalizeConfig(c, paths)

	if len(out.Targets) != 0 {
		t.Fatalf("expected 0 targets, got %d", len(out.Targets))
	}
}

func TestNormalizeConfig_CommandsUntouched(t *testing.T) {
	c := Config{
		Targets: []string{"~/file"},
		Commands: []string{
			"echo ~",
			"cd ~",
		},
	}

	paths := env.Paths{
		Home: "/home/test",
	}

	out := normalizeConfig(c, paths)

	if out.Commands[0] != "echo ~" || out.Commands[1] != "cd ~" {
		t.Errorf("commands should not be modified, got %+v", out.Commands)
	}
}
