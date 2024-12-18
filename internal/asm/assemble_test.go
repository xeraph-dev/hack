// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssemble(t *testing.T) {
	asm := `@123
			@LABEL
			@test
			(LABEL)
			@456`

	prog, _ := Parse(strings.NewReader(asm))

	str := strings.Builder{}
	assert.Nil(t, prog.Assemble(&str))
	assert.Equal(t, `0000000001111011
0000000000000100
0000000000010001
0000000111001000`, str.String())
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
