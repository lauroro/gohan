// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package cli

import (
	"fmt"
)

// printHelp prints the global CLI help.
//
// It shows:
// - program description
// - usage syntax
// - available commands
// - positional arguments (derived from commands)
func printHelp() {
	fmt.Println("USAGE:")
	fmt.Println("  gohan <command> [arg]")
	fmt.Println()

	fmt.Println("COMMANDS:")

	maxLen := maxUsageLen()

	for _, c := range commands {
		fmt.Printf("  %-*s  %s\n", maxLen, c.Usage(), c.Description)
	}

	printPositionalArgs()
}

// printCommandHelp prints detailed help for a single command.
//
// It includes:
// - program description
// - command usage
// - command description
// - positional argument documentation (if required)
func printCommandHelp(c Command) {
	fmt.Printf("USAGE:\n  gohan %s\n", c.Usage())
	fmt.Println()

	fmt.Println(c.Description)

	if c.RequiredArg {
		fmt.Println()
		fmt.Println("POSITIONAL ARGUMENTS:")
		fmt.Printf("  %-15s %s\n", c.ArgName, c.ArgDescription)
	}
}

// Usage returns command usage string including argument if required.
func (c Command) Usage() string {
	if c.RequiredArg {
		return fmt.Sprintf("%s <%s>", c.Name, c.ArgName)
	}
	return c.Name
}

// maxUsageLen returns the longest formatted command usage string.
// It is used to align CLI help output.
func maxUsageLen() int {
	max := 0
	for _, c := range commands {
		if l := len(c.Usage()); l > max {
			max = l
		}
	}
	return max
}

// collectPositionalArgs returns a deduplicated list of all positional args.
func collectPositionalArgs() map[string]string {
	args := make(map[string]string)

	for _, c := range commands {
		if c.RequiredArg {
			args[c.ArgName] = c.ArgDescription
		}
	}

	return args
}

// printPositionalArgs prints all positional arguments used by commands.
func printPositionalArgs() {
	args := collectPositionalArgs()

	if len(args) == 0 {
		return
	}

	fmt.Println()
	fmt.Println("POSITIONAL ARGUMENTS:")

	for name, desc := range args {
		fmt.Printf("  %-15s %s\n", name, desc)
	}
}
