// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

type Config struct {
	Targets  []string `yaml:"targets"`
	Commands []string `yaml:"commands"`
	Features Features `yaml:"features"`
}

type Features struct {
	GTKAutoColorScheme bool `yaml:"gtk_auto_colorscheme"`
}
