// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package env

import (
	"os"
	"testing"
)

func TestEnv_WithExistingVariable(t *testing.T) {
	key := "TEST_ENV_VAR"
	expected := "value"

	os.Setenv(key, expected)
	defer os.Unsetenv(key)

	result := env(key, "fallback")

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestEnv_WithFallback(t *testing.T) {
	key := "NON_EXISTING_ENV_VAR"

	os.Unsetenv(key)

	result := env(key, "fallback")

	if result != "fallback" {
		t.Errorf("expected fallback, got %s", result)
	}
}

func TestEnv_PanicWhenMissingAndNoFallback(t *testing.T) {
	key := "ANOTHER_NON_EXISTING_ENV_VAR"

	os.Unsetenv(key)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, but did not panic")
		}
	}()

	env(key, "")
}
