// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePushStatement(t *testing.T) {
	line := "push constant 5"

	stmt, err := ParseStatement(line)
	assert.Nil(t, err)
	assert.Equal(t, Statement{
		Command: CommandPush,
		Segment: SegmentConstant,
		Index:   5,
	}, stmt)
}

func TestParsePopStatement(t *testing.T) {
	line := "pop constant 5"

	stmt, err := ParseStatement(line)
	assert.Nil(t, err)
	assert.Equal(t, Statement{
		Command: CommandPop,
		Segment: SegmentConstant,
		Index:   5,
	}, stmt)
}

func TestParseAddStatement(t *testing.T) {
	line := "add"

	stmt, err := ParseStatement(line)
	assert.Nil(t, err)
	assert.Equal(t, Statement{Command: CommandAdd}, stmt)
}
