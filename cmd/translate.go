// Copyright 2024 xeraph. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"hack/internal/vm"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

var translateCommand = &cobra.Command{
	Use:  "translate",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmFilePath := args[0]
		asmFilePath := strings.TrimSuffix(vmFilePath, path.Ext(vmFilePath)) + ".asm"

		if prog, err := parseVM(vmFilePath); err != nil {
			log.Fatal(err)
		} else if err = translate(asmFilePath, prog); err != nil {
			log.Fatal(err)
		}
	},
}

func parseVM(filePath string) (prog vm.Program, err error) {
	var file *os.File
	if file, err = os.Open(filePath); err != nil {
		return
	}
	defer file.Close()

	if prog, err = vm.Parse(file); err != nil {
		return
	}

	return
}

func translate(filePath string, prog vm.Program) (err error) {
	if prog == nil {
		panic("prog is nil")
	}

	var file *os.File
	if file, err = os.Create(filePath); err != nil {
		return
	}
	defer file.Close()

	if err = prog.Translate(file); err != nil {
		return
	}

	return
}
