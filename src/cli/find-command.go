// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package cli

// findCommand searches for a command by name.
//
// It returns a pointer to the matching Command if found,
// or nil if no command with the given name exists.
func findCommand(name string) *Command {
	for i := range commands {
		if commands[i].Name == name {
			return &commands[i]
		}
	}
	return nil
}
