// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import "testing"

func TestParseConfig_Valid(t *testing.T) {
	input := `
targets:
  - file1.txt
  - file2.txt
commands:
  - cmd1
  - cmd2
`

	cfg, err := parseConfig(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cfg.Targets) != 2 {
		t.Fatalf("expected 2 targets, got %d", len(cfg.Targets))
	}

	if len(cfg.Commands) != 2 {
		t.Fatalf("expected 2 commands, got %d", len(cfg.Commands))
	}

	if cfg.Targets[0] != "file1.txt" || cfg.Targets[1] != "file2.txt" {
		t.Errorf("unexpected targets: %#v", cfg.Targets)
	}

	if cfg.Commands[0] != "cmd1" || cfg.Commands[1] != "cmd2" {
		t.Errorf("unexpected commands: %#v", cfg.Commands)
	}
}

func TestParseConfig_Empty(t *testing.T) {
	input := ``

	cfg, err := parseConfig(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cfg.Targets) != 0 {
		t.Errorf("expected empty targets, got %#v", cfg.Targets)
	}

	if len(cfg.Commands) != 0 {
		t.Errorf("expected empty commands, got %#v", cfg.Commands)
	}
}

func TestParseConfig_MissingFields(t *testing.T) {
	input := `
targets:
  - only-target
`

	cfg, err := parseConfig(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cfg.Targets) != 1 {
		t.Errorf("expected 1 target, got %d", len(cfg.Targets))
	}

	if len(cfg.Commands) != 0 {
		t.Errorf("expected empty commands, got %#v", cfg.Commands)
	}
}

func TestParseConfig_InvalidYAML(t *testing.T) {
	input := `
targets:
  - ok
  - bad: [
`

	_, err := parseConfig(input)

	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestParseConfig_OrderPreserved(t *testing.T) {
	input := `
targets:
  - a
  - b
  - c
commands:
  - x
  - y
`

	cfg, err := parseConfig(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedTargets := []string{"a", "b", "c"}
	expectedCommands := []string{"x", "y"}

	for i, v := range expectedTargets {
		if cfg.Targets[i] != v {
			t.Errorf("targets order mismatch at %d: expected %s, got %s", i, v, cfg.Targets[i])
		}
	}

	for i, v := range expectedCommands {
		if cfg.Commands[i] != v {
			t.Errorf("commands order mismatch at %d: expected %s, got %s", i, v, cfg.Commands[i])
		}
	}
}
