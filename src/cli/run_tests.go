// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package cli

import (
	"errors"
	"testing"
)

func TestRun_NoArgs(t *testing.T) {
	commands = []Command{}
	if err := Run([]string{}); err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func TestRun_GlobalHelp(t *testing.T) {
	commands = []Command{}
	if err := Run([]string{"--help"}); err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func TestRun_UnknownCommand(t *testing.T) {
	commands = []Command{}
	if err := Run([]string{"unknown"}); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestRun_CommandHelp(t *testing.T) {
	commands = []Command{
		{Name: "apply"},
	}
	if err := Run([]string{"apply", "--help"}); err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func TestRun_CommandWithArg(t *testing.T) {
	called := false
	commands = []Command{
		{
			Name:        "apply",
			RequiredArg: true,
			Run: func(arg string) error {
				called = true
				if arg != "test" {
					t.Errorf("expected arg 'test', got %s", arg)
				}
				return nil
			},
		},
	}

	if err := Run([]string{"apply", "test"}); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !called {
		t.Errorf("expected command to be executed")
	}
}

func TestRun_MissingRequiredArg(t *testing.T) {
	commands = []Command{
		{
			Name:        "apply",
			RequiredArg: true,
			ArgName:     "palette",
			Run: func(arg string) error {
				return nil
			},
		},
	}

	if err := Run([]string{"apply"}); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestRun_CommandReturnsError(t *testing.T) {
	commands = []Command{
		{
			Name: "apply",
			Run: func(arg string) error {
				return errors.New("fail")
			},
		},
	}

	if err := Run([]string{"apply"}); err == nil {
		t.Fatal("expected error, got nil")
	}
}
