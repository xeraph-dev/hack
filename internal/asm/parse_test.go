// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	asm := `@123
			@LABEL
			@test
			(LABEL)
			@456`

	prog, err := Parse(strings.NewReader(asm))
	assert.Nil(t, err)
	assert.Equal(t, Program{
		&AddressInstructionConstant{Address: 123},
		&AddressInstructionSymbol{Symbol: "LABEL"},
		&AddressInstructionSymbol{Symbol: "test"},
		&LabelInstruction{Symbol: "LABEL"},
		&AddressInstructionConstant{Address: 456},
	}, prog)
}

func TestParseAddressInstructionConstant(t *testing.T) {
	line := "@123"

	instr, err := ParseAddressInstruction(line)
	assert.Nil(t, err)
	assert.Equal(t, &AddressInstructionConstant{Address: 123}, instr)
}

func TestParseAddressInstructionSymbol(t *testing.T) {
	line := "@test"

	instr, err := ParseAddressInstruction(line)
	assert.Nil(t, err)
	assert.Equal(t, &AddressInstructionSymbol{Symbol: "test"}, instr)
}

func TestParseAddressInstructionInvalid(t *testing.T) {
	line := "@1invalid"

	_, err := ParseAddressInstruction(line)
	assert.ErrorIs(t, err, ErrAddressInstructionInvalid)
}

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
