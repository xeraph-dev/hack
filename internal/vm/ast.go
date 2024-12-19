// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vm

type (
	Command int

	Segment int

	Statement struct {
		Command Command
		Segment Segment
		Index   int16
	}
)
