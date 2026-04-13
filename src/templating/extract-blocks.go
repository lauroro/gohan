// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import "strings"

const (
	startMarker = "@gohan:start"
	endMarker   = "@gohan:end"
)

// extractBlocks finds all gohan blocks in the given content.
// It returns all blocks with their positions and inner content.
func extractBlocks(content string) []block {
	var blocks []block

	i := 0
	for {
		start := strings.Index(content[i:], startMarker)
		if start == -1 {
			break
		}
		start += i

		end := strings.Index(content[start:], endMarker)
		if end == -1 {
			// ignore unclosed blocks
			break
		}
		end += start

		innerStart := start + len(startMarker)
		innerEnd := end

		blocks = append(blocks, block{
			Start:   start,
			End:     end + len(endMarker),
			Content: content[innerStart:innerEnd],
		})

		i = end + len(endMarker)
	}

	return blocks
}
