// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
