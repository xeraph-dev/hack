// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestAssembleAddressInstruction(t *testing.T) {
	line := "@123"

	instr, err := ParseAddressInstruction(line)
	assert.Nil(t, err)
	assert.Equal(t, &AddressInstructionConstant{Address: 123}, instr)

	str := strings.Builder{}
	assert.Nil(t, instr.Assemble(&str))

	assert.Equal(t, "0000000001111011", str.String())
}
