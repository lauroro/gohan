// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package cli

import "fmt"

// printHelp prints the global CLI usage information.
//
// It displays the general usage format and the list of
// available commands with their short descriptions.
func printHelp() {
	fmt.Println("USAGE:")
	fmt.Println("  gohan <command> [arg]")
	fmt.Println()
	fmt.Println("COMMANDS:")
	for _, c := range commands {
		fmt.Printf("  %-10s %s\n", c.Name, c.Description)
	}
}

// printCommandHelp prints detailed usage information for a specific command.
//
// It shows the command usage, including required arguments if any,
// followed by the command's description.
func printCommandHelp(c Command) {
	fmt.Printf("USAGE:\n  gohan %s", c.Name)
	if c.RequiredArg {
		fmt.Printf(" <%s>", c.ArgName)
	}
	fmt.Println()
	fmt.Println(c.Description)
}
