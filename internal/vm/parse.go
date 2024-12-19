// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vm

import (
	"strconv"
	"strings"
)

func ParseStatement(line string) (stmt Statement, err error) {
	var before, after string
	var found bool

	before, after, found = strings.Cut(line, " ")

	if stmt.Command, err = ParseCommand(before); !found || err != nil {
		return
	}

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
	stmt.Index = int16(index)

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
