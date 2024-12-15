// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLabelInstruction(t *testing.T) {
	line := "(LABEL)"

	instr, err := ParseLabelInstruction(line)
	assert.Nil(t, err)
	assert.Equal(t, &LabelInstruction{Symbol: "LABEL"}, instr)
}

func TestParseLabelInstructionInvalid(t *testing.T) {
	line := "(1INVALID)"

	_, err := ParseLabelInstruction(line)
	assert.ErrorIs(t, err, ErrLabelInstructionInvalid)
}
