// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	prog, err := ParseString(`
@123
@LABEL
@test
(LABEL)
@456
A=M;JMP
`)
	assert.Nil(t, err)
	assert.Equal(t, Program{
		&AddressInstructionConstant{Address: 123},
		&AddressInstructionSymbol{Symbol: "LABEL"},
		&AddressInstructionSymbol{Symbol: "test"},
		&LabelInstruction{Symbol: "LABEL"},
		&AddressInstructionConstant{Address: 456},
		&ComputeInstruction{Dest: DestA, Comp: Comp1M, Jump: JumpJMP},
	}, prog)
}

func TestParseAddressInstruction(t *testing.T) {
	t.Run("constant", func(t *testing.T) {
		instr, err := ParseAddressInstruction("@123")
		assert.Nil(t, err)
		assert.Equal(t, &AddressInstructionConstant{Address: 123}, instr)
	})

	t.Run("symbol", func(t *testing.T) {
		instr, err := ParseAddressInstruction("@test")
		assert.Nil(t, err)
		assert.Equal(t, &AddressInstructionSymbol{Symbol: "test"}, instr)
	})
}

func TestParseLabelInstruction(t *testing.T) {
	instr, err := ParseLabelInstruction("(LABEL)")
	assert.Nil(t, err)
	assert.Equal(t, &LabelInstruction{Symbol: "LABEL"}, instr)
}

func TestParseComputeInstruction(t *testing.T) {
	t.Run("comp", func(t *testing.T) {
		instr, err := ParseComputeInstruction("0")
		assert.Nil(t, err)
		assert.Equal(t, &ComputeInstruction{Comp: Comp00}, instr)
	})

	t.Run("dest+comp", func(t *testing.T) {
		instr, err := ParseComputeInstruction("D=M")
		assert.Nil(t, err)
		assert.Equal(t, &ComputeInstruction{Dest: DestD, Comp: Comp1M}, instr)
	})

	t.Run("comp+jump", func(t *testing.T) {
		instr, err := ParseComputeInstruction("M;JMP")
		assert.Nil(t, err)
		assert.Equal(t, &ComputeInstruction{Comp: Comp1M, Jump: JumpJMP}, instr)
	})

	t.Run("dest+comp+jump", func(t *testing.T) {
		instr, err := ParseComputeInstruction("A=M;JMP")
		assert.Nil(t, err)
		assert.Equal(t, &ComputeInstruction{Dest: DestA, Comp: Comp1M, Jump: JumpJMP}, instr)
	})
}
