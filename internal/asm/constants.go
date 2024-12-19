// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import "regexp"

const (
	SymbolR0  = "R0"
	SymbolR1  = "R1"
	SymbolR2  = "R2"
	SymbolR3  = "R3"
	SymbolR4  = "R4"
	SymbolR5  = "R5"
	SymbolR6  = "R6"
	SymbolR7  = "R7"
	SymbolR8  = "R8"
	SymbolR9  = "R9"
	SymbolR10 = "R10"
	SymbolR11 = "R11"
	SymbolR12 = "R12"
	SymbolR13 = "R13"
	SymbolR14 = "R14"
	SymbolR15 = "R15"

	SymbolSP   = "SP"
	SymbolLCL  = "LCL"
	SymbolARG  = "ARG"
	SymbolTHIS = "THIS"
	SymbolTHAT = "THAT"

	SymbolSCREEN = "SCREEN"
	SymbolKBD    = "KBD"
)

var (
	AddressRegex = regexp.MustCompile("^[0-9]+$")
	SymbolRegex  = regexp.MustCompile("^[a-zA-Z_.$:][0-9a-zA-Z_.$:]*$")

	DefaultSymbols = map[string]int16{
		SymbolR0:  0,
		SymbolR1:  1,
		SymbolR2:  2,
		SymbolR3:  3,
		SymbolR4:  4,
		SymbolR5:  5,
		SymbolR6:  6,
		SymbolR7:  7,
		SymbolR8:  8,
		SymbolR9:  9,
		SymbolR10: 10,
		SymbolR11: 11,
		SymbolR12: 12,
		SymbolR13: 13,
		SymbolR14: 14,
		SymbolR15: 15,

		SymbolSP:   0,
		SymbolLCL:  1,
		SymbolARG:  2,
		SymbolTHIS: 3,
		SymbolTHAT: 4,

		SymbolSCREEN: 16384,
		SymbolKBD:    24576,
	}
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

	CompToString = map[Comp]string{
		Comp00:       "0",
		Comp01:       "1",
		Comp0Neg1:    "-1",
		Comp0D:       "D",
		Comp0A:       "A",
		Comp0NotD:    "!D",
		Comp0NotA:    "!A",
		Comp0NegD:    "-D",
		Comp0NegA:    "-A",
		Comp0APlus1:  "A+1",
		Comp0DPlus1:  "D+1",
		Comp0DMinus1: "D-1",
		Comp0AMinus1: "A-1",
		Comp0DPlusA:  "D+A",
		Comp0AMinusD: "A-D",
		Comp0DAndA:   "D&A",
		Comp0DOrA:    "D|A",
		Comp1M:       "M",
		Comp1NotM:    "!M",
		Comp1NegM:    "-M",
		Comp1MPlus1:  "M+1",
		Comp1MMinus1: "M-1",
		Comp1DPlusM:  "D+M",
		Comp1DMinusM: "D-M",
		Comp1MMinusD: "M-D",
		Comp1DAndM:   "D&M",
		Comp1DOrM:    "D|M",
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

	JumpToString = map[Jump]string{
		JumpNull: "",
		JumpJGT:  "JGT",
		JumpJEQ:  "JEQ",
		JumpJGE:  "JGE",
		JumpJLT:  "JLT",
		JumpJNE:  "JNE",
		JumpJLE:  "JLE",
		JumpJMP:  "JMP",
	}
)
