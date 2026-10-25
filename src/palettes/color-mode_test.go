// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package palettes

import "testing"

func TestGetColorMode(t *testing.T) {
	tests := []struct {
		name    string
		palette Palette
		want    string
	}{
		{
			name: "dark",
			palette: Palette{
				Base00: "#181a1f",
				Base01: "#262b33",
				Base02: "#303843",
				Base03: "#4f565e",
				Base04: "#7b8c9a",
				Base05: "#a9b7c2",
				Base06: "#c3d0d8",
				Base07: "#e0ebf2",
			},
			want: "dark",
		},
		{
			name: "light",
			palette: Palette{
				Base00: "#fbeef2",
				Base01: "#f6dfe4",
				Base02: "#dcc7cc",
				Base03: "#725e63",
				Base04: "#645053",
				Base05: "#534244",
				Base06: "#3f3738",
				Base07: "#2f2627",
			},
			want: "light",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetColorMode(tt.palette)
			if got != tt.want {
				t.Errorf("expected %s, got %s", tt.want, got)
			}
		})
	}
}
