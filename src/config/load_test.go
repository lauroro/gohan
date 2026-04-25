// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import (
	"os"
	"path/filepath"
	"testing"

	"gohan/src/env"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestLoadConfig_Success(t *testing.T) {
	content := `
targets:
  - "~/file1"
  - "~/file2"
commands:
  - "echo hello"
`

	dir := t.TempDir()

	file := filepath.Join(dir, "config.yaml")
	writeFile(t, file, content)

	paths := env.Paths{
		ConfigFile: file,
		Home:       "/home/test",
	}

	cfg, err := LoadConfig(paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cfg.Targets) != 2 {
		t.Fatalf("expected 2 targets, got %d", len(cfg.Targets))
	}

	if cfg.Targets[0] != "/home/test/file1" {
		t.Errorf("expected normalized path, got %s", cfg.Targets[0])
	}

	if len(cfg.Commands) != 1 {
		t.Errorf("expected 1 command, got %d", len(cfg.Commands))
	}
}

func TestLoadConfig_FileMissing(t *testing.T) {
	paths := env.Paths{
		ConfigFile: filepath.Join(t.TempDir(), "missing.yaml"),
	}

	_, err := LoadConfig(paths)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestLoadConfig_InvalidYAML(t *testing.T) {
	dir := t.TempDir()

	file := filepath.Join(dir, "config.yaml")
	writeFile(t, file, `: invalid yaml`)

	paths := env.Paths{
		ConfigFile: file,
	}

	_, err := LoadConfig(paths)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestLoadConfig_InvalidStructure(t *testing.T) {
	dir := t.TempDir()

	file := filepath.Join(dir, "config.yaml")
	writeFile(t, file, `
targets: "not-an-array"
`)

	paths := env.Paths{
		ConfigFile: file,
	}

	_, err := LoadConfig(paths)
	if err == nil {
		t.Fatal("expected validation error but got nil")
	}
}

func TestLoadConfig_Normalization(t *testing.T) {
	dir := t.TempDir()

	file := filepath.Join(dir, "config.yaml")
	writeFile(t, file, `
targets:
  - "~/test"
`)

	paths := env.Paths{
		ConfigFile: file,
		Home:       "/home/user",
	}

	cfg, err := LoadConfig(paths)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Targets[0] != "/home/user/test" {
		t.Errorf("expected expanded path, got %s", cfg.Targets[0])
	}
}
