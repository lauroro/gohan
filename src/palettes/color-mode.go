// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import (
	"strconv"
	"strings"
)

// GetColorMode classifies a palette as "dark" or "light" based on the
// average relative luminance of its Base00, Base01 and Base02 colors.
// It return "dark" when the average is below 0.5, and "light" otherwise.
func GetColorMode(p Palette) string {
	colors := []string{
		p.Base00, p.Base01, p.Base02,
	}

	var total float64

	for _, color := range colors {
		color = strings.TrimPrefix(color, "#")
		value, _ := strconv.ParseUint(color, 16, 32)

		r := float64((value>>16)&0xff) / 255
		g := float64((value>>8)&0xff) / 255
		b := float64(value&0xff) / 255

		total += relativeLuminance(r)*0.2126 +
			relativeLuminance(g)*0.7152 +
			relativeLuminance(b)*0.0722
	}

	if total/float64(len(colors)) < 0.5 {
		return "dark"
	}

	return "light"
}

func relativeLuminance(channel float64) float64 {
	if channel <= 0.04045 {
		return channel / 12.92
	}

	adjusted := (channel + 0.055) / 1.055
	return adjusted * adjusted * adjusted
}
