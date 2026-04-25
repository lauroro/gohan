// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import "testing"

func TestExecCmd_Success(t *testing.T) {
	out, err := ExecCmd("echo hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out == "" {
		t.Errorf("expected output, got empty string")
	}
}

func TestExecCmd_Failure(t *testing.T) {
	_, err := ExecCmd("exit 1")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestExecCmd_CombinedOutput(t *testing.T) {
	out, err := ExecCmd("echo out && echo err 1>&2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !(contains(out, "out") && contains(out, "err")) {
		t.Errorf("expected combined output, got %q", out)
	}
}

func TestExecCmd_Pipe(t *testing.T) {
	out, err := ExecCmd("echo hello | grep hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !contains(out, "hello") {
		t.Errorf("expected 'hello' in output, got %q", out)
	}
}

func TestExecCmd_EmptyCommand(t *testing.T) {
	out, err := ExecCmd("")
	if err != nil {
		t.Fatalf("unexpected error for empty command: %v", err)
	}

	if out != "" {
		t.Errorf("expected empty output, got %q", out)
	}
}

func contains(s, substr string) bool {
	return indexOf(s, substr) >= 0
}

func indexOf(s, substr string) int {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
