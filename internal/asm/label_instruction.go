// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"errors"
	"io"
	"strings"
)

var (
	ErrLabelInstructionInvalid = errors.New("invalid label instruction")
)

type LabelInstruction struct {
	Symbol string
}

func (instr *LabelInstruction) Assemble(io.Writer) (err error) {
	panic("LabelInstruction cannot be assembled")
}

func ParseLabelInstruction(line string) (instr Instruction, err error) {
	symbol := strings.TrimPrefix(line, "(")
	symbol = strings.TrimSuffix(symbol, ")")

	if SymbolRegex.MatchString(symbol) {
		instr = &LabelInstruction{Symbol: symbol}
	} else {
		err = ErrLabelInstructionInvalid
	}

	return
}
