// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

// Block represents a @gohan templated section inside a file.
//
// Start and End are byte indices referring to the original content string.
// Content contains the inner text between the markers (excluding them).
type block struct {
	Start   int
	End     int
	Content string
}
