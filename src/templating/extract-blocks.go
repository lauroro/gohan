// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

const (
	startMarker = "@gohan:start"
	endMarker   = "@gohan:end"
)

// extractBlocks finds all gohan blocks in the given content.
// It returns all blocks with their positions and inner content.
func extractBlocks(content string) []block
