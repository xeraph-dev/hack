// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"bufio"
	"io"
	"regexp"
	"slices"
	"strings"
)

var (
	AddressRegex = regexp.MustCompile("^[0-9]+$")
	SymbolRegex  = regexp.MustCompile("^[a-zA-Z_.$:][0-9a-zA-Z_.$:]*$")

	DefaultSymbols = map[string]int16{
		"R0":  0,
		"R1":  1,
		"R2":  2,
		"R3":  3,
		"R4":  4,
		"R5":  5,
		"R6":  6,
		"R7":  7,
		"R8":  8,
		"R9":  9,
		"R10": 10,
		"R11": 11,
		"R12": 12,
		"R13": 13,
		"R14": 14,
		"R15": 15,

		"SP":   0,
		"LCL":  1,
		"ARG":  2,
		"THIS": 3,
		"THAT": 4,

		"SCREEN": 16384,
		"KBD":    24576,
	}
)

type (
	Instruction interface {
		Assemble(io.Writer) error
	}

	Program []Instruction
)

func ParseString(str string) (prog Program, err error) {
	return Parse(strings.NewReader(str))
}

func Parse(r io.Reader) (prog Program, err error) {
	var line string
	var instr Instruction

	s := bufio.NewScanner(r)

	for s.Scan() {
		line = strings.Trim(s.Text(), " \t\r")

		switch {
		case line == "" || strings.HasPrefix(line, "//"):
			continue
		case strings.HasPrefix(line, "@"):
			if instr, err = ParseAddressInstruction(line); err != nil {
				return
			}
		case strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")"):
			if instr, err = ParseLabelInstruction(line); err != nil {
				return
			}
		default:
			if instr, err = ParseComputeInstruction(line); err != nil {
				return
			}
		}

		prog = append(prog, instr)
	}

	err = s.Err()
	return
}

func (prog Program) AssembleString() (str string, err error) {
	builder := strings.Builder{}
	err = prog.Assemble(&builder)
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
