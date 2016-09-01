package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path"
)

// execute go run goshell.go, if failed, deprecated these codes
func execute() {
	flush()
	err := tryExecute()
	if err != nil {
		undo()
	} else {
		save()
	}
}

// try goimports to remove unused imports and add requried imports
// try go run goshell.go, if fail, return err
func tryExecute() error {
	// run goimports in advance
	goimports := exec.Command("goimports", "-w", path.Join(codeDir, codeFile))
	goimports.Run()

	// run go file
	cmd := exec.Command("go", "run", path.Join(codeDir, codeFile))

	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer
	err := cmd.Run()
	fmt.Print(outBuffer.String())
	fmt.Print(errBuffer.String())

	if errBuffer.Len() != 0 {
		err = errors.New("run err")
	}
	return err
}
