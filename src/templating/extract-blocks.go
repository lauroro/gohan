// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import (
	"strings"
)

const (
	startMarker = "@gohan:start"
	endMarker   = "@gohan:end"
)

// extractBlocks scans the input content and extracts all @gohan template blocks.
//
// A block is defined as any section starting with a line containing
// "@gohan:start" and ending with a line containing "@gohan:end".
//
// The function is tolerant to additional characters before or after the markers
// (e.g. comments like "// @gohan:start" or "/* @gohan:end */").
//
// Behavior:
//   - Only the first occurrence of a start marker begins a block.
//   - The block ends at the first following end marker.
//   - Markers themselves are not included in the content.
//   - Empty or whitespace-only blocks are ignored.
//   - Content preserve internal newlines but excludes markers.
func extractBlocks(content string) []block {
	var blocks []block

	lines := strings.Split(content, "\n")

	var (
		inBlock   bool
		startByte int
		cursor    int
		buffer    strings.Builder
	)

	for _, line := range lines {
		lineLen := len(line) + 1

		if !inBlock {
			if strings.Contains(line, startMarker) {
				inBlock = true

				startOffset := strings.Index(line, startMarker)
				startByte = cursor + startOffset
				buffer.Reset()
			}
		} else {
			if strings.Contains(line, endMarker) {
				endOffset := strings.Index(line, endMarker)
				endByte := cursor + endOffset + len(endMarker)

				content := buffer.String()

				if len(strings.TrimSpace(content)) > 0 {
					blocks = append(blocks, block{
						Start:   startByte,
						End:     endByte,
						Content: content,
					})
				}

				inBlock = false
				buffer.Reset()

				cursor += lineLen
				continue
			}

			buffer.WriteString(line)
			buffer.WriteByte('\n')
		}

		cursor += lineLen
	}

	return blocks
}
