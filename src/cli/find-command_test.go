// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package cli

import "testing"

func TestFindCommand_Found(t *testing.T) {
	commands = []Command{
		{Name: "apply"},
		{Name: "list"},
	}

	cmd := findCommand("list")

	if cmd == nil {
		t.Fatal("expected command, got nil")
	}

	if cmd.Name != "list" {
		t.Errorf("expected list, got %s", cmd.Name)
	}
}

func TestFindCommand_NotFound(t *testing.T) {
	commands = []Command{
		{Name: "apply"},
	}

	cmd := findCommand("unknown")

	if cmd != nil {
		t.Errorf("expected nil, got %+v", cmd)
	}
}

func TestFindCommand_ReturnsPointerToOriginal(t *testing.T) {
	commands = []Command{
		{Name: "apply"},
	}

	cmd := findCommand("apply")

	if cmd == nil {
		t.Fatal("expected command, got nil")
	}

	cmd.Name = "modified"

	if commands[0].Name != "modified" {
		t.Errorf("expected modification to affect original slice")
	}
}
