// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"hack/internal/asm"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

var assembleCommand = &cobra.Command{
	Use:  "assemble",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		asmFilePath := args[0]
		hackFilePath := strings.TrimSuffix(asmFilePath, path.Ext(asmFilePath)) + ".hack"

		if prog, err := parse(asmFilePath); err != nil {
			log.Fatal(err)
		} else if err = assemble(hackFilePath, prog); err != nil {
			log.Fatal(err)
		}
	},
}

func parse(filePath string) (prog asm.Program, err error) {
	var file *os.File
	if file, err = os.Open(filePath); err != nil {
		return
	}
	defer file.Close()

	if prog, err = asm.Parse(file); err != nil {
		return
	}

	return
}

func assemble(filePath string, prog asm.Program) (err error) {
	if prog == nil {
		panic("prog is nil")
	}

	var file *os.File
	if file, err = os.Create(filePath); err != nil {
		return
	}
	defer file.Close()

	if err = prog.Assemble(file); err != nil {
		return
	}

	return
}
