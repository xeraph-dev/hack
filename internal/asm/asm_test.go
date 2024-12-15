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

func TestResolveSymbols(t *testing.T) {
	asm := `@123
			@LABEL
			@test
			(LABEL)
			@456`

	prog, _ := Parse(strings.NewReader(asm))
	prog.ResolveSymbols()
	assert.Equal(t, Program{
		&AddressInstructionConstant{Address: 123},
		&AddressInstructionConstant{Address: 4},
		&AddressInstructionConstant{Address: 17},
		&AddressInstructionConstant{Address: 456},
	}, prog)
}

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
