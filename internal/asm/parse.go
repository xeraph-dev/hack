// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"bufio"
	"io"
	"strconv"
	"strings"
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

func ParseAddressInstruction(line string) (instr Instruction, err error) {
	addrOrSym := strings.TrimPrefix(line, "@")

	if AddressRegex.MatchString(addrOrSym) {
		var address int
		if address, err = strconv.Atoi(addrOrSym); err != nil {
			return
		}
		instr = &AddressInstructionConstant{Address: int16(address)}
	} else if SymbolRegex.MatchString(addrOrSym) {
		instr = &AddressInstructionSymbol{Symbol: addrOrSym}
	} else {
		err = ErrAddressInstructionInvalid
	}

	return
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

func ParseComputeInstruction(line string) (instr Instruction, err error) {
	computeInstruction := &ComputeInstruction{}

	var before, after string
	var found bool
	if before, after, found = strings.Cut(line, "="); !found {
		after = before
		before = ""
	}
	if computeInstruction.Dest, err = ParseComputeInstructionDest(before); err != nil {
		return
	}
	before, after, _ = strings.Cut(after, ";")
	if computeInstruction.Comp, err = ParseComputeInstructionComp(before); err != nil {
		return
	}
	if computeInstruction.Jump, err = ParseComputeInstructionJump(after); err != nil {
		return
	}

	instr = computeInstruction
	return
}

func ParseComputeInstructionComp(str string) (comp Comp, err error) {
	var ok bool
	if comp, ok = StringToComp[str]; !ok {
		err = ErrCompInvalid{comp: str}
	}
	return
}

func ParseComputeInstructionDest(str string) (dest Dest, err error) {
	for _, char := range str {
		switch char {
		case 'A':
			dest |= DestA
		case 'D':
			dest |= DestD
		case 'M':
			dest |= DestM
		}
	}
	return
}

func ParseComputeInstructionJump(str string) (jump Jump, err error) {
	var ok bool
	if jump, ok = StringToJump[str]; !ok {
		err = ErrJumpInvalid{jump: str}
	}
	return
}
