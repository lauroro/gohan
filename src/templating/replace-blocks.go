// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import (
	"fmt"
	"strings"
)

// replaceBlocks reconstructs the file by replacing the inner content of each block,
// while preserving markers and all surrounding text.
//
// It assumes:
// - blocks are ordered by Start
// - Start/End are valid byte indices into original
// - Content contains the new inner content (already transformed)
func replaceBlocks(original string, blocks []block) (string, error) {
	var result strings.Builder
	last := 0

	for _, b := range blocks {
		if b.Start < last || b.End > len(original) {
			return "", fmt.Errorf("invalid block range")
		}

		result.WriteString(original[last:b.Start])

		full := original[b.Start:b.End]

		startIdx := strings.Index(full, startMarker)
		endIdx := strings.Index(full, endMarker)

		if startIdx == -1 || endIdx == -1 {
			return "", fmt.Errorf("invalid block markers")
		}

		startIdx += len(startMarker)

		if startIdx < len(full) && full[startIdx] == '\n' {
			startIdx++
		}

		result.WriteString(
			full[:startIdx] +
				b.Content +
				full[endIdx:],
		)

		last = b.End
	}

	result.WriteString(original[last:])

	return result.String(), nil
}
