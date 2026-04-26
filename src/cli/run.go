// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package cli

import "fmt"

// Run parses CLI arguments and dispatches the corresponding command.
//
// It supports:
// - version printing via "--version"
// - global help via "-h" or "--help"
// - command lookup by name
// - command-specific help via "<command> --help"
// - optional or required single argument per command
//
// If no arguments are provided, the global help is printed.
//
// It returns an error if:
// - the command is unknown
// - a required argument is missing
// - the command execution return an error
func Run(args []string) error {
	if len(args) == 0 {
		printHelp()
		return nil
	}

	if args[0] == "--version" {
		version()
		return nil
	}

	if args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return nil
	}

	name := args[0]

	cmd := findCommand(name)
	if cmd == nil {
		return fmt.Errorf("unknown command: %s", name)
	}

	if len(args) > 1 && (args[1] == "--help" || args[1] == "-h") {
		printCommandHelp(*cmd)
		return nil
	}

	var arg string
	if len(args) > 1 {
		arg = args[1]
	}

	if cmd.RequiredArg && arg == "" {
		return fmt.Errorf("%s requires %s", cmd.Name, cmd.ArgName)
	}

	return cmd.Run(arg)
}
