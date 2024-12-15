// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	ErrAddressInstructionInvalid = errors.New("invalid address instruction")
)

type (
	AddressInstructionConstant struct {
		Address int16
	}

	AddressInstructionSymbol struct {
		Symbol string
	}
)

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
