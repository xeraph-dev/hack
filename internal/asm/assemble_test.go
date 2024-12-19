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
	prog, _ := ParseString(`
@123
@LABEL
@test
(LABEL)
@456
A=M;JMP
`)

	bin, err := AssembleString(prog)
	assert.Nil(t, err)
	assert.Equal(t, strings.Trim(`
0000000001111011
0000000000000011
0000000000010000
0000000111001000
1111110000100111
`, " \t\n\r"), bin)
}

func TestAssembleAddressInstruction(t *testing.T) {
	instr, err := ParseAddressInstruction("@123")
	assert.Nil(t, err)
	assert.Equal(t, &AddressInstructionConstant{Address: 123}, instr)

	bin, err := AssembleString(instr)
	assert.Nil(t, err)

	assert.Equal(t, "0000000001111011", bin)
}
