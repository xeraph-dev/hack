// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import "errors"

var (
	ErrAddressInstructionInvalid = errors.New("invalid address instruction")
	ErrLabelInstructionInvalid   = errors.New("invalid label instruction")
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
