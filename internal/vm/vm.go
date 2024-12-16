// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vm

import (
	"bufio"
	"io"
	"strings"
)

type (
	Program []Statement
)

func ParseString(str string) (prog Program, err error) {
	return Parse(strings.NewReader(str))
}

func Parse(r io.Reader) (prog Program, err error) {
	var line string
	var stmt Statement

	s := bufio.NewScanner(r)

	for s.Scan() {
		line = strings.Trim(s.Text(), " \t\r")
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		if stmt, err = ParseStatement(line); err != nil {
			return
		}

		prog = append(prog, stmt)
	}

	err = s.Err()
	return
}

func (prog Program) TranslateString() (str string, err error) {
	builder := strings.Builder{}
	err = prog.Translate(&builder)
	str = builder.String()
	return
}

func (prog Program) Translate(w io.Writer) (err error) {
	for idx, stmt := range prog {
		if idx > 0 {
			if _, err = w.Write([]byte{'\n'}); err != nil {
				return
			}
		}
		stmt.Translate(w)
	}
	return
}
