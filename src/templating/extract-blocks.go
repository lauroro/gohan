// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import (
	"strings"
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

	offset := 0

	for {
		start := strings.Index(content[offset:], startMarker)
		if start == -1 {
			break
		}
		start += offset

		startLineEnd := strings.Index(content[start:], "\n")
		if startLineEnd == -1 {
			break
		}
		startLineEnd += start + 1

		end := strings.Index(content[startLineEnd:], endMarker)
		if end == -1 {
			break
		}
		end += startLineEnd

		endLineStart := strings.LastIndex(content[:end], "\n")
		if endLineStart == -1 {
			endLineStart = 0
		} else {
			endLineStart += 1
		}

		inner := content[startLineEnd:endLineStart]

		if len(strings.TrimSpace(inner)) > 0 {
			blocks = append(blocks, block{
				Start:   start,
				End:     end + len(endMarker),
				Content: inner,
			})
		}

		offset = end + len(endMarker)
	}

	return blocks
}
