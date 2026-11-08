// SPDX-FileCopyrightText: 2026 lauroro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import "testing"

func validConfig() Config {
	return Config{
		Targets:  []string{"a", "b"},
		Commands: []string{"cmd1", "cmd2"},
	}
}

func TestValidateConfig_Valid(t *testing.T) {
	if err := validateConfig(validConfig()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateConfig_EmptyTargets(t *testing.T) {
	c := Config{
		Targets: []string{},
	}

	if err := validateConfig(c); err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestValidateConfig_EmptyTargetValue(t *testing.T) {
	c := Config{
		Targets: []string{"ok", ""},
	}

	if err := validateConfig(c); err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestValidateConfig_EmptyCommandValue(t *testing.T) {
	c := Config{
		Targets:  []string{"ok"},
		Commands: []string{"cmd", ""},
	}

	if err := validateConfig(c); err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestValidateConfig_NoCommands(t *testing.T) {
	c := Config{
		Targets: []string{"ok"},
	}

	if err := validateConfig(c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateConfig_FeaturesEnabled(t *testing.T) {
	c := Config{
		Targets: []string{"ok"},
		Features: Features{
			GTKAutoColorScheme: true,
		},
	}

	if err := validateConfig(c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateConfig_FeaturesDisabledByDefault(t *testing.T) {
	c := Config{
		Targets: []string{"ok"},
	}

	if c.Features.GTKAutoColorScheme {
		t.Fatal("expected gtk_auto_colorscheme to default to false")
	}

	if err := validateConfig(c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
