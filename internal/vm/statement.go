// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vm

import (
	"io"
	"strconv"
	"strings"
)

type (
	ErrStatementInvalid struct {
		stmt string
	}

	ErrCommandInvalid struct {
		cmd string
	}

	ErrSegmentInvalid struct {
		seg string
	}
)

func (err ErrStatementInvalid) Error() string {
	return "invalid statement: " + err.stmt
}

func (err ErrCommandInvalid) Error() string {
	return "invalid command: " + err.cmd
}

func (err ErrSegmentInvalid) Error() string {
	return "invalid segment: " + err.seg
}

type (
	Command int

	Segment int

	Statement struct {
		Command Command
		Segment Segment
		Index   uint16
	}
)

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

func ParseStatement(line string) (stmt Statement, err error) {
	var before, after string
	var found bool

	if before, after, found = strings.Cut(line, " "); !found {
		err = ErrStatementInvalid{stmt: line}
		return
	}

	if stmt.Command, err = ParseCommand(before); err != nil {
		return
	}

	if after != "" {
		if before, after, found = strings.Cut(after, " "); !found {
			err = ErrStatementInvalid{stmt: line}
			return
		}

		if stmt.Segment, err = ParseSegment(before); err != nil {
			return
		}

		var index uint64
		if index, err = strconv.ParseUint(after, 10, 16); err != nil {
			return
		}
		stmt.Index = uint16(index)
	}

	return
}

func ParseCommand(str string) (cmd Command, err error) {
	var ok bool
	if cmd, ok = StringToCommand[str]; !ok {
		err = ErrCommandInvalid{cmd: str}
	}
	return
}

func ParseSegment(str string) (seg Segment, err error) {
	var ok bool
	if seg, ok = StringToSegment[str]; !ok {
		err = ErrSegmentInvalid{seg: str}
	}
	return
}

func (stmt Statement) TranslateString() (str string, err error) {
	builder := strings.Builder{}
	err = stmt.Translate(&builder)
	str = builder.String()
	return
}

func (stmt Statement) Translate(w io.Writer) (err error) {
	return
}