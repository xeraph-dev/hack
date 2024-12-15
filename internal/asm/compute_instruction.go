// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"io"
	"strconv"
	"strings"
)

type (
	ErrCompInvalid struct {
		comp string
	}
	ErrDestInvalid struct {
		dest string
	}
	ErrJumpInvalid struct {
		jump string
	}
)

func (err ErrCompInvalid) Error() string {
	return "invalid comp: " + err.comp
}

func (err ErrDestInvalid) Error() string {
	return "invalid dest: " + err.dest
}

func (err ErrJumpInvalid) Error() string {
	return "invalid jump: " + err.jump
}

type (
	ComputeInstruction struct {
		Comp Comp
		Dest Dest
		Jump Jump
	}

	Comp interface {
		Instruction
		A() uint8
	}
	Comp0 int
	Comp1 int

	Dest int

	Jump int
)

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

const (
	DestNull Dest = iota
	DestM
	DestD
	DestDM
	DestA
	DestAM
	DestAD
	DestADM
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

var (
	StringToComp = map[string]Comp{
		"0":   Comp00,
		"1":   Comp01,
		"-1":  Comp0Neg1,
		"D":   Comp0D,
		"A":   Comp0A,
		"!D":  Comp0NotD,
		"!A":  Comp0NotA,
		"-D":  Comp0NegD,
		"-A":  Comp0NegA,
		"A+1": Comp0APlus1,
		"D+1": Comp0DPlus1,
		"D-1": Comp0DMinus1,
		"A-1": Comp0AMinus1,
		"D+A": Comp0DPlusA,
		"A-D": Comp0AMinusD,
		"D&A": Comp0DAndA,
		"D|A": Comp0DOrA,

		"M":   Comp1M,
		"!M":  Comp1NotM,
		"-M":  Comp1NegM,
		"M+1": Comp1MPlus1,
		"M-1": Comp1MMinus1,
		"D+M": Comp1DPlusM,
		"D-M": Comp1DMinusM,
		"M-D": Comp1MMinusD,
		"D&M": Comp1DAndM,
		"D|M": Comp1DOrM,
	}

	StringToDest = map[string]Dest{
		"":    DestNull,
		"M":   DestM,
		"D":   DestD,
		"DM":  DestDM,
		"MD":  DestDM,
		"A":   DestA,
		"AM":  DestAM,
		"AD":  DestAD,
		"ADM": DestADM,
	}

	StringToJump = map[string]Jump{
		"":    JumpNull,
		"JGT": JumpJGT,
		"JEQ": JumpJEQ,
		"JGE": JumpJGE,
		"JLT": JumpJLT,
		"JNE": JumpJNE,
		"JLE": JumpJLE,
		"JMP": JumpJMP,
	}
)

func ParseComputeInstruction(line string) (instr Instruction, err error) {
	computeInstruction := &ComputeInstruction{}

	var before, after string
	var found bool
	if before, after, found = strings.Cut(line, "="); !found {
		after = before
		before = ""
	}
	if computeInstruction.Dest, err = ParseDest(before); err != nil {
		return
	}
	before, after, _ = strings.Cut(after, ";")
	if computeInstruction.Comp, err = ParseComp(before); err != nil {
		return
	}
	if computeInstruction.Jump, err = ParseJump(after); err != nil {
		return
	}

	instr = computeInstruction
	return
}

func ParseComp(str string) (comp Comp, err error) {
	var ok bool
	if comp, ok = StringToComp[str]; !ok {
		err = ErrCompInvalid{comp: str}
	}
	return
}

func ParseDest(str string) (dest Dest, err error) {
	var ok bool
	if dest, ok = StringToDest[str]; !ok {
		err = ErrDestInvalid{dest: str}
	}
	return
}

func ParseJump(str string) (jump Jump, err error) {
	var ok bool
	if jump, ok = StringToJump[str]; !ok {
		err = ErrJumpInvalid{jump: str}
	}
	return
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
	switch dest {
	case DestNull, DestM, DestD, DestDM, DestA, DestAM, DestAD, DestADM:
		bin := strconv.FormatInt(int64(dest), 2)
		pad := strings.Repeat("0", 3-len(bin))
		if _, err = w.Write([]byte(pad)); err != nil {
			return
		}
		if _, err = w.Write([]byte(bin)); err != nil {
			return
		}
	default:
		panic("unhandled Dest: " + strconv.Itoa(int(dest)))
	}

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

func (comp Comp0) A() uint8 {
	return 0
}

func (comp Comp1) A() uint8 {
	return 1
}
