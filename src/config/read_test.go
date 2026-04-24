// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package config

import (
	"os"
	"testing"
)

func TestReadConfigFile_Success(t *testing.T) {
	tmp := t.TempDir()
	file := tmp + "/config.yaml"

	err := os.WriteFile(file, []byte("hello"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	result, err := readConfigFile(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != "hello" {
		t.Errorf("expected hello, got %s", result)
	}
}

func TestReadConfigFile_NotFound(t *testing.T) {
	_, err := readConfigFile("/non/existent/file")
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
