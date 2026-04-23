// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package cli

import "fmt"

var commands = []Command{
	{
		Name:           "apply",
		Description:    "Apply a palette",
		ArgName:        "palette",
		ArgDescription: "palette name, file name without extension",
		RequiredArg:    true,
		Run: func(arg string) error {
			fmt.Println("apply:", arg)
			return nil
		},
	},
	{
		Name:        "list",
		Description: "List palettes",
		Run: func(arg string) error {
			fmt.Println("list")
			return nil
		},
	},
}
