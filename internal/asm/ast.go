// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"io"
	"slices"
)

type (
	Assemblable interface {
		Assemble(io.Writer) error
	}

	Formattable interface {
		Format(io.Writer) error
	}

	Instruction interface {
		Assemblable
		Formattable
		__instruction()
	}

	Program []Instruction

	AddressInstructionConstant struct {
		Address int16
	}

	AddressInstructionSymbol struct {
		Symbol string
	}

	LabelInstruction struct {
		Symbol string
	}

	ComputeInstruction struct {
		Comp Comp
		Dest Dest
		Jump Jump
	}

	Comp interface {
		Assemblable
		Formattable
		A() uint8
	}
	Comp0 int
	Comp1 int

	Dest int

	Jump int
)

func (instr *AddressInstructionConstant) __instruction() {}
func (instr *AddressInstructionSymbol) __instruction()   {}
func (instr *LabelInstruction) __instruction()           {}
func (instr *ComputeInstruction) __instruction()         {}

func (comp Comp0) A() uint8 {
	return 0
}

func (comp Comp1) A() uint8 {
	return 1
}

func (prog *Program) ResolveSymbols() {
	syms := DefaultSymbols
	var labels []Instruction

	var line int16 = 0
	for _, instr := range *prog {
		switch instr.(type) {
		case *LabelInstruction:
			labelInstr := instr.(*LabelInstruction)
			syms[labelInstr.Symbol] = line
			labels = append(labels, labelInstr)
		default:
			line += 1
		}
	}

	var address int16 = 16
	for i, instr := range *prog {
		switch instr.(type) {
		case *AddressInstructionSymbol:
			addrInstrSym := instr.(*AddressInstructionSymbol)
			if _, ok := syms[addrInstrSym.Symbol]; !ok {
				syms[addrInstrSym.Symbol] = address
				address += 1
			}
			(*prog)[i] = &AddressInstructionConstant{Address: syms[addrInstrSym.Symbol]}
		}
	}

	*prog = slices.DeleteFunc(*prog, func(instr Instruction) bool {
		return slices.Contains(labels, instr)
	})

	return
}

const (
	Comp00 Comp0 = iota
	Comp01
	Comp0Neg1
	Comp0D
	Comp0A
	Comp0NotD
	Comp0NotA
	Comp0NegD
	Comp0NegA
	Comp0DPlus1
	Comp0APlus1
	Comp0DMinus1
	Comp0AMinus1
	Comp0DPlusA
	Comp0DMinusA
	Comp0AMinusD
	Comp0DAndA
	Comp0DOrA
)

const (
	Comp1M Comp1 = iota
	Comp1NotM
	Comp1NegM
	Comp1MPlus1
	Comp1MMinus1
	Comp1DPlusM
	Comp1DMinusM
	Comp1MMinusD
	Comp1DAndM
	Comp1DOrM
)

const DestNull = 0
const (
	DestA Dest = 1 << iota
	DestD
	DestM
)

const (
	JumpNull Jump = iota
	JumpJGT
	JumpJEQ
	JumpJGE
	JumpJLT
	JumpJNE
	JumpJLE
	JumpJMP
)
