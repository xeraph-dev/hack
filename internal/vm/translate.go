// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vm

import (
	"hack/internal/asm"
	"io"
	"strconv"
	"strings"
)

func (stmt Statement) TranslateString() (str string, err error) {
	builder := strings.Builder{}
	err = stmt.Translate(&builder)
	str = builder.String()
	return
}

func (stmt Statement) Translate(w io.Writer) (err error) {
	switch stmt.Command {
	case CommandPush:
		switch stmt.Segment {
		case SegmentConstant:
			asm.Program{
				&asm.AddressInstructionConstant{Address: stmt.Index},
				&asm.ComputeInstruction{Dest: asm.DestD, Comp: asm.Comp0A},
				&asm.AddressInstructionSymbol{Symbol: asm.SymbolSP},
				&asm.ComputeInstruction{Dest: asm.DestA, Comp: asm.Comp1M},
				&asm.ComputeInstruction{Dest: asm.DestM, Comp: asm.Comp0D},
				&asm.AddressInstructionSymbol{Symbol: asm.SymbolSP},
				&asm.ComputeInstruction{Dest: asm.DestM, Comp: asm.Comp1MPlus1},
			}.Format(w)
		default:
			panic("unhandled segment for CommandPush: " + strconv.Itoa(int(stmt.Segment)))
		}
	case CommandAdd:
		asm.Program{
			&asm.AddressInstructionSymbol{Symbol: asm.SymbolSP},
			&asm.ComputeInstruction{Dest: asm.DestA | asm.DestM, Comp: asm.Comp1MMinus1},
			&asm.ComputeInstruction{Dest: asm.DestD, Comp: asm.Comp1M},
			&asm.AddressInstructionSymbol{Symbol: asm.SymbolSP},
			&asm.ComputeInstruction{Dest: asm.DestA, Comp: asm.Comp1MMinus1},
			&asm.ComputeInstruction{Dest: asm.DestM, Comp: asm.Comp1DPlusM},
		}.Format(w)
	default:
		panic("unhandled command: " + strconv.Itoa(int(stmt.Command)))
	}

	return
}
