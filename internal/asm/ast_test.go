// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveSymbols(t *testing.T) {
	prog, _ := ParseString(`
@123
@LABEL
@test
(LABEL)
@456
A=M;JMP
`)
	prog.ResolveSymbols()
	assert.Equal(t, Program{
		&AddressInstructionConstant{Address: 123},
		&AddressInstructionConstant{Address: 3},
		&AddressInstructionConstant{Address: 16},
		&AddressInstructionConstant{Address: 456},
		&ComputeInstruction{Dest: DestA, Comp: Comp1M, Jump: JumpJMP},
	}, prog)
}
