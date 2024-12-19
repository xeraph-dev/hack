// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vm

const (
	CommandPush Command = iota
	CommandPop

	CommandAdd
	CommandSub
	CommandNeg

	CommandEq
	CommandGt
	CommandLt

	CommandAnd
	CommandOr
	CommandNot
)

const (
	SegmentNull Segment = iota
	SegmentArgument
	SegmentLocal
	SegmentStatic
	SegmentConstant
	SegmentThis
	SegmentThat
	SegmentPointer
	SegmentTemp
)

var (
	StringToCommand = map[string]Command{
		"push": CommandPush,
		"pop":  CommandPop,

		"add": CommandAdd,
		"sub": CommandSub,
		"neg": CommandNeg,

		"eq": CommandEq,
		"gt": CommandGt,
		"lt": CommandLt,

		"and": CommandAnd,
		"or":  CommandOr,
		"not": CommandNot,
	}

	StringToSegment = map[string]Segment{
		"argument": SegmentArgument,
		"local":    SegmentLocal,
		"static":   SegmentStatic,
		"constant": SegmentConstant,
		"this":     SegmentThis,
		"that":     SegmentThat,
		"pointer":  SegmentPointer,
		"temp":     SegmentTemp,
	}
)
