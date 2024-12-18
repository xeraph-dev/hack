// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import "io"

func FormatString(expr Formattable) (str string, err error) {
	return
}

func (prog Program) Format(w io.Writer) (err error) {
	return
}

func (instr *AddressInstructionConstant) Format(w io.Writer) (err error) {
	return
}

func (instr *AddressInstructionSymbol) Format(w io.Writer) (err error) {
	return
}

func (instr *LabelInstruction) Format(w io.Writer) (err error) {
	return
}

func (instr *ComputeInstruction) Format(w io.Writer) (err error) {
	return
}

func (comp Comp0) Format(w io.Writer) (err error) {
	return
}

func (comp Comp1) Format(w io.Writer) (err error) {
	return
}

func (dest Dest) Format(w io.Writer) (err error) {
	return
}

func (jump Jump) Format(w io.Writer) (err error) {
	return
}
