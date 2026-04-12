// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package env

import "os"

// env retrieves the value of the environment variable identified by `key`.
// If the variable is set, its value is returned.
// If it is not set and a non-empty `fallback` is provided, the fallback is returned.
// If the variable is not set and no fallback is given, the function panics.
func env(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else if fallback != "" {
		return fallback
	} else {
		panic("cannot provide a value for env var: " + key)
	}
}
