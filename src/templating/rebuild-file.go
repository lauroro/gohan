// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import (
	"fmt"
	"strings"
)

// rebuildFile reconstructs the file by replacing only the inner content of each block,
// while preserving all markers and external content.
//
// It assumes that block.Start and block.End refer to the original string byte indices
// and include both markers.
func rebuildFile(original string, blocks []block, newContents []string) (string, error) {
	if len(blocks) != len(newContents) {
		return "", fmt.Errorf("blocks and newContents length mismatch")
	}

	var result strings.Builder
	last := 0

	for i, b := range blocks {
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

		rebuilt :=
			full[:startIdx] +
				newContents[i] +
				full[endIdx:]

		result.WriteString(rebuilt)

		last = b.End
	}

	result.WriteString(original[last:])
	return result.String(), nil
}
