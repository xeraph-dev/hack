// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"io"
	"strconv"
	"strings"
)

func AssembleString(expr Assemblable) (str string, err error) {
	builder := strings.Builder{}
	err = expr.Assemble(&builder)
	str = builder.String()
	return
}

func (prog Program) Assemble(w io.Writer) (err error) {
	prog.ResolveSymbols()
	for idx, instr := range prog {
		if idx > 0 {
			if _, err = w.Write([]byte{'\n'}); err != nil {
				return
			}
		}
		instr.Assemble(w)
	}
	return
}

func (instr *AddressInstructionConstant) Assemble(w io.Writer) (err error) {
	bin := strconv.FormatInt(int64(instr.Address), 2)
	pad := strings.Repeat("0", 15-len(bin))

	if _, err = w.Write([]byte{'0'}); err != nil {
		return
	}
	if _, err = w.Write([]byte(pad)); err != nil {
		return
	}
	if _, err = w.Write([]byte(bin)); err != nil {
		return
	}

	return
}

func (instr *AddressInstructionSymbol) Assemble(io.Writer) error {
	panic("AddressInstructionSymbol cannot be assembled")
}

func (instr *LabelInstruction) Assemble(io.Writer) (err error) {
	panic("LabelInstruction cannot be assembled")
}

func (instr *ComputeInstruction) Assemble(w io.Writer) (err error) {
	if _, err = w.Write([]byte{'1', '1', '1'}); err != nil {
		return
	}

	if instr.Comp == nil {
		panic("ComputeInstruction.Comp is nil")
	}
	if _, err = w.Write([]byte(strconv.FormatUint(uint64(instr.Comp.A()), 2))); err != nil {
		return
	}

	if err = instr.Comp.Assemble(w); err != nil {
		return
	}

	if err = instr.Dest.Assemble(w); err != nil {
		return
	}

	if err = instr.Jump.Assemble(w); err != nil {
		return
	}

	return
}

func (comp Comp0) Assemble(w io.Writer) (err error) {
	switch comp {
	case Comp00:
		_, err = w.Write([]byte("101010"))
	case Comp01:
		_, err = w.Write([]byte("111111"))
	case Comp0Neg1:
		_, err = w.Write([]byte("111010"))
	case Comp0D:
		_, err = w.Write([]byte("001100"))
	case Comp0A:
		_, err = w.Write([]byte("110000"))
	case Comp0NotD:
		_, err = w.Write([]byte("001101"))
	case Comp0NotA:
		_, err = w.Write([]byte("110001"))
	case Comp0NegD:
		_, err = w.Write([]byte("001111"))
	case Comp0NegA:
		_, err = w.Write([]byte("110011"))
	case Comp0DPlus1:
		_, err = w.Write([]byte("011111"))
	case Comp0APlus1:
		_, err = w.Write([]byte("110111"))
	case Comp0DMinus1:
		_, err = w.Write([]byte("001110"))
	case Comp0AMinus1:
		_, err = w.Write([]byte("110010"))
	case Comp0DPlusA:
		_, err = w.Write([]byte("000010"))
	case Comp0DMinusA:
		_, err = w.Write([]byte("010011"))
	case Comp0AMinusD:
		_, err = w.Write([]byte("000111"))
	case Comp0DAndA:
		_, err = w.Write([]byte("000000"))
	case Comp0DOrA:
		_, err = w.Write([]byte("010101"))
	default:
		panic("unhandled Comp0: " + strconv.Itoa(int(comp)))
	}

	return
}

func (comp Comp1) Assemble(w io.Writer) (err error) {
	switch comp {
	case Comp1M:
		_, err = w.Write([]byte("110000"))
	case Comp1NotM:
		_, err = w.Write([]byte("110001"))
	case Comp1NegM:
		_, err = w.Write([]byte("110011"))
	case Comp1MPlus1:
		_, err = w.Write([]byte("110111"))
	case Comp1MMinus1:
		_, err = w.Write([]byte("110010"))
	case Comp1DPlusM:
		_, err = w.Write([]byte("000010"))
	case Comp1DMinusM:
		_, err = w.Write([]byte("010011"))
	case Comp1MMinusD:
		_, err = w.Write([]byte("000111"))
	case Comp1DAndM:
		_, err = w.Write([]byte("000000"))
	case Comp1DOrM:
		_, err = w.Write([]byte("010101"))
	default:
		panic("unhandled Comp1: " + strconv.Itoa(int(comp)))
	}

	return
}

func (dest Dest) Assemble(w io.Writer) (err error) {
	var a, d, m byte
	if dest&DestA != 0 {
		a = '1'
	} else {
		a = '0'
	}
	if dest&DestD != 0 {
		d = '1'
	} else {
		d = '0'
	}
	if dest&DestM != 0 {
		m = '1'
	} else {
		m = '0'
	}
	w.Write([]byte{a, d, m})
	return
}

func (jump Jump) Assemble(w io.Writer) (err error) {
	switch jump {
	case JumpNull, JumpJGT, JumpJEQ, JumpJGE, JumpJLT, JumpJNE, JumpJLE, JumpJMP:
		bin := strconv.FormatInt(int64(jump), 2)
		pad := strings.Repeat("0", 3-len(bin))
		if _, err = w.Write([]byte(pad)); err != nil {
			return
		}
		if _, err = w.Write([]byte(bin)); err != nil {
			return
		}
	default:
		panic("unhandled Jump: " + strconv.Itoa(int(jump)))
	}

	return
}
