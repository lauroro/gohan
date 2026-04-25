// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import "testing"

func TestRebuildFile_SingleBlock(t *testing.T) {
	input := `before
@gohan:start 
hello
@gohan:end
after`

	blocks := extractBlocks(input)

	blocks[0].Content = "NEW\n"

	out, err := rebuildFile(input, blocks)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !contains(out, "NEW\n") {
		t.Errorf("expected NEW content, got:\n%s", out)
	}

	if !contains(out, "@gohan:start") || !contains(out, "@gohan:end") {
		t.Errorf("markers should be preserved")
	}
}

func TestRebuildFile_MultipleBlocks(t *testing.T) {
	input := `@gohan:start
a
@gohan:end
middle
@gohan:start
b
@gohan:end`

	blocks := extractBlocks(input)

	blocks[0].Content = "X\n"
	blocks[1].Content = "Y\n"

	out, err := rebuildFile(input, blocks)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !contains(out, "X\n") || !contains(out, "Y\n") {
		t.Errorf("blocks not replaced correctly:\n%s", out)
	}
}

func TestRebuildFile_PreservesOutsideText(t *testing.T) {
	input := `AAA
@gohan:start
x
@gohan:end
BBB`

	blocks := extractBlocks(input)

	blocks[0].Content = "Z\n"

	out, err := rebuildFile(input, blocks)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !contains(out, "AAA") || !contains(out, "BBB") {
		t.Errorf("outside text not preserved:\n%s", out)
	}
}

func TestRebuildFile_InvalidRange(t *testing.T) {
	b := block{
		Start:   0,
		End:     9999,
		Content: "x",
	}

	_, err := rebuildFile("short", []block{b})
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestRebuildFile_InvalidMarkers(t *testing.T) {
	input := `no markers here`

	blocks := []block{
		{
			Start:   0,
			End:     len(input),
			Content: "x",
		},
	}

	_, err := rebuildFile(input, blocks)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
