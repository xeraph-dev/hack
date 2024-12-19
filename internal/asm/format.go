// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"io"
	"strconv"
)

func FormatString(expr Formattable) (str string, err error) {
	return
}

func (prog Program) Format(w io.Writer) (err error) {
	for idx, instr := range prog {
		if idx > 0 {
			if _, err = w.Write([]byte{'\n'}); err != nil {
				return
			}
		}
		instr.Format(w)
	}
	return
}

func (instr *AddressInstructionConstant) Format(w io.Writer) (err error) {
	if _, err = w.Write([]byte{'@'}); err != nil {
		return
	}
	if _, err = w.Write([]byte(strconv.Itoa(int(instr.Address)))); err != nil {
		return
	}
	return
}

func (instr *AddressInstructionSymbol) Format(w io.Writer) (err error) {
	if _, err = w.Write([]byte{'@'}); err != nil {
		return
	}
	if _, err = w.Write([]byte(instr.Symbol)); err != nil {
		return
	}
	return
}

func (instr *LabelInstruction) Format(w io.Writer) (err error) {
	if _, err = w.Write([]byte{'('}); err != nil {
		return
	}
	if _, err = w.Write([]byte(instr.Symbol)); err != nil {
		return
	}
	if _, err = w.Write([]byte{')'}); err != nil {
		return
	}
	return
}

func (instr *ComputeInstruction) Format(w io.Writer) (err error) {
	instr.Dest.Format(w)
	instr.Comp.Format(w)
	instr.Jump.Format(w)
	return
}

func (comp Comp0) Format(w io.Writer) (err error) {
	if _, err = w.Write([]byte(CompToString[comp])); err != nil {
		return
	}
	return
}

func (comp Comp1) Format(w io.Writer) (err error) {
	if _, err = w.Write([]byte(CompToString[comp])); err != nil {
		return
	}
	return
}

func (dest Dest) Format(w io.Writer) (err error) {
	if dest == DestNull {
		return
	}

	var d []byte
	if dest&DestA != 0 {
		d = append(d, 'A')
	}
	if dest&DestD != 0 {
		d = append(d, 'D')
	}
	if dest&DestM != 0 {
		d = append(d, 'M')
	}
	d = append(d, '=')
	if _, err = w.Write(d); err != nil {
		return
	}
	return
}

func (jump Jump) Format(w io.Writer) (err error) {
	if jump == JumpNull {
		return
	}

	if _, err = w.Write([]byte{';'}); err != nil {
		return
	}
	if _, err = w.Write([]byte(JumpToString[jump])); err != nil {
		return
	}
	return
}
