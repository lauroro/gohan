// SPDX-FileCopyrightText: 2026 roro
// SPDX-License-Identifier: GPL-3.0-or-later

package templating

import (
	"fmt"

	"gohan/src/palettes"
	"gohan/src/tools"
)

// TransformFile applies a palette transformation to a target file.
//
// It performs the full pipeline:
// - reads the file from disk
// - extracts templated blocks
// - transforms their content using the provided palettes
// - rebuilds the file preserving markers and surrounding content
// - writes the result back to disk
//
// It returns an error if any step of the process fails.
func TransformFile(target string, old, new palettes.Palette) error {
	content, err := tools.ReadFile(target)
	if err != nil {
		return fmt.Errorf("could not transform target %s: %w", target, err)
	}

	blocks := extractBlocks(content)

	blocks, err = transformBlocks(blocks, old, new)
	if err != nil {
		return fmt.Errorf("could not transform target %s: %w", target, err)
	}

	newContent, err := replaceBlocks(content, blocks)
	if err != nil {
		return fmt.Errorf("could not transform target %s: %w", target, err)
	}

	if err := tools.WriteFile(target, newContent); err != nil {
		return fmt.Errorf("could not transform target %s: %w", target, err)
	}

	return nil
}
