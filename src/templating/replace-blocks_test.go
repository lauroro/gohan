// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import "testing"

func TestReplaceBlocks_SingleBlock(t *testing.T) {
	input := "before\n@gohan:start\nhello\n@gohan:end\nafter"

	blocks := []block{
		{
			Start:   7,
			End:     7 + len("@gohan:start\nhello\n@gohan:end"),
			Content: "NEW\n",
		},
	}

	out, err := replaceBlocks(input, blocks)
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

func TestReplaceBlocks_MultipleBlocks(t *testing.T) {
	input := "@gohan:start\na\n@gohan:end\nmiddle\n@gohan:start\nb\n@gohan:end"

	block1 := "@gohan:start\na\n@gohan:end"
	block2 := "@gohan:start\nb\n@gohan:end"

	blocks := []block{
		{
			Start:   0,
			End:     len(block1),
			Content: "X\n",
		},
		{
			Start:   len(block1) + len("\nmiddle\n"),
			End:     len(block1) + len("\nmiddle\n") + len(block2),
			Content: "Y\n",
		},
	}

	out, err := replaceBlocks(input, blocks)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !contains(out, "X\n") || !contains(out, "Y\n") {
		t.Errorf("blocks not replaced correctly:\n%s", out)
	}
}

func TestReplaceBlocks_PreservesOutsideText(t *testing.T) {
	input := "AAA\n@gohan:start\nx\n@gohan:end\nBBB"

	blockStr := "@gohan:start\nx\n@gohan:end"

	blocks := []block{
		{
			Start:   4,
			End:     4 + len(blockStr),
			Content: "Z\n",
		},
	}

	out, err := replaceBlocks(input, blocks)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !contains(out, "AAA") || !contains(out, "BBB") {
		t.Errorf("outside text not preserved:\n%s", out)
	}
}

func TestReplaceBlocks_InvalidRange(t *testing.T) {
	b := block{
		Start:   0,
		End:     9999,
		Content: "x",
	}

	_, err := replaceBlocks("short", []block{b})
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestReplaceBlocks_InvalidMarkers(t *testing.T) {
	input := "no markers here"

	blocks := []block{
		{
			Start:   0,
			End:     len(input),
			Content: "x",
		},
	}

	_, err := replaceBlocks(input, blocks)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
