// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import "testing"

func TestExtractBlocks_Single(t *testing.T) {
	input := `before
@gohan:start
hello
@gohan:end
after`

	blocks := extractBlocks(input)

	if len(blocks) != 1 {
		t.Fatalf("expected 1 block, got %d", len(blocks))
	}

	expected := "\nhello\n"

	if blocks[0].Content != expected {
		t.Errorf("unexpected content: %q", blocks[0].Content)
	}
}

func TestExtractBlocks_Multiple(t *testing.T) {
	input := `@gohan:start
a
@gohan:end
middle
@gohan:start
b
@gohan:end`

	blocks := extractBlocks(input)

	if len(blocks) != 2 {
		t.Fatalf("expected 2 blocks, got %d", len(blocks))
	}

	if blocks[0].Content != "\na\n" {
		t.Errorf("unexpected first block: %q", blocks[0].Content)
	}

	if blocks[1].Content != "\nb\n" {
		t.Errorf("unexpected second block: %q", blocks[1].Content)
	}
}

func TestExtractBlocks_None(t *testing.T) {
	input := `no blocks here`

	blocks := extractBlocks(input)

	if len(blocks) != 0 {
		t.Fatalf("expected 0 blocks, got %d", len(blocks))
	}
}

func TestExtractBlocks_Unclosed(t *testing.T) {
	input := `@gohan:start
hello`

	blocks := extractBlocks(input)

	if len(blocks) != 0 {
		t.Fatalf("expected 0 blocks, got %d", len(blocks))
	}
}

func TestExtractBlocks_EmptyBlock(t *testing.T) {
	input := `@gohan:start@gohan:end`

	blocks := extractBlocks(input)

	if len(blocks) != 1 {
		t.Fatalf("expected 1 block, got %d", len(blocks))
	}

	if blocks[0].Content != "" {
		t.Errorf("expected empty content, got %q", blocks[0].Content)
	}
}

func TestExtractBlocks_PreservesIndices(t *testing.T) {
	input := `xx@gohan:start
abc
@gohan:endyy`

	blocks := extractBlocks(input)

	if len(blocks) != 1 {
		t.Fatalf("expected 1 block, got %d", len(blocks))
	}

	b := blocks[0]

	reconstructed := input[b.Start:b.End]

	expected := `@gohan:start
abc
@gohan:end`

	if reconstructed != expected {
		t.Errorf("expected %q, got %q", expected, reconstructed)
	}
}
