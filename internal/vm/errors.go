// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vm

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
