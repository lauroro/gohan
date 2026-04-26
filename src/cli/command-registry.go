// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package cli

import (
	"fmt"
	"gohan/src/actions"
	"gohan/src/env"
)

var commands = []Command{
	{
		Name:           "apply",
		Description:    "Apply a palette",
		ArgName:        "palette",
		ArgDescription: "palette name, file name without extension",
		RequiredArg:    true,
		Run: func(arg string) error {
			return actions.Apply(arg, env.BuildPaths())
		},
	},
	{
		Name:        "list",
		Description: "List palettes",
		Run: func(arg string) error {
			list, err := actions.List(env.BuildPaths())
			if err != nil {
				return err
			}

			for _, p := range list {
				fmt.Println(p)
			}

			return nil
		},
	},
	{
		Name:           "init",
		Description:    "Initialize gohan state",
		ArgName:        "palette",
		ArgDescription: "palette name, file name without extension",
		RequiredArg:    true,
		Run: func(arg string) error {
			return actions.Init(arg, env.BuildPaths())
		},
	},
}
