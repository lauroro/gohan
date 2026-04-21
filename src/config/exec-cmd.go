// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import (
	"fmt"
	"os/exec"
)

// execCmd executes a shell command string using the system shell.
//
// It runs the command via "sh -c", so it supports pipes, redirects,
// and compound commands (e.g. "echo hi && ls").
//
// It returns the combined stdout and stderr output.
// If the command exits with a non-zero status, the error includes
// the command output for easier debugging.
func execCmd(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command failed: %s: %w\noutput:\n%s", command, err, string(out))
	}

	return string(out), nil
}
