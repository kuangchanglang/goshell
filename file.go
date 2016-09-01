package main

import (
	"os"
	"path"
)

// in case import conflict while import "os" manually
const (
	importOS      = "import GOSHELL_OS \"os\"\n"
	importSyscall = "import GOSHELL_SYSCALL \"syscall\"\n"
)

// mask stdout to nil
const funcMaskStdout = `
func _MASK_STDOUT(){
	GOSHELL_OS.Stdout = nil
}
`
const funcUnmarskStdout = `
func _UNMASK_STDOUT(){
	GOSHELL_OS.Stdout = GOSHELL_OS.NewFile(uintptr(GOSHELL_SYSCALL.Stdout), "/dev/stdout")
}
`

// buffers that will be flush to file and execute go run
// we do flush() and gorun() for each non-assignment and
// non-func statement
var (
	importBuffer *Buffer // import codes
	funcBuffer   *Buffer // func codes
	mainBuffer   *Buffer // codes in main function
	increBuffer  *Buffer // used to record increment codes that has not been run
)

// init buffer
func initBuffer() {
	importBuffer = NewBufferFromString("")
	funcBuffer = NewBufferFromString("")
	mainBuffer = NewBufferFromString("")
	increBuffer = NewBufferFromString("")

	importBuffer.WriteString("package main\n")
	importBuffer.WriteString(importOS)
	importBuffer.WriteString(importSyscall)
	funcBuffer.WriteString(funcMaskStdout)
	funcBuffer.WriteString(funcUnmarskStdout)
	mainBuffer.WriteString("func main(){\n")
	mainBuffer.WriteString("_MASK_STDOUT()\n") // mask stdout at the beginning then unmask before new codes
	save()
}

// setup env
func setUp() error {
	initBuffer()
	err := os.MkdirAll(codeDir, permission)
	return err
}

// clean up env
func cleanUp() {
	os.Remove(path.Join(codeDir, codeFile))
	// remove dir?
}

// flush buffer to file, overwrite file if exists
// write buffer in such order: import , func, main
// main section contains three part: old main,
// unmask stdout and increment buffer
func flush() error {
	file, err := os.Create(path.Join(codeDir, codeFile))
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(importBuffer.String())
	file.WriteString(funcBuffer.String())
	file.WriteString(mainBuffer.String())
	file.WriteString("_UNMASK_STDOUT()\n")
	file.WriteString(increBuffer.String())
	// write } to end of func main
	file.WriteString("}")
	return nil
}

// save buffer if go run return no errors
// append increment buffer to main and reset it
func save() {
	braceStack = 0
	importBuffer.Save()
	funcBuffer.Save()
	// main buffer should append increment buffer
	mainBuffer.Write(increBuffer.Bytes())
	increBuffer.Reset()
	mainBuffer.Save()
}

// undo buffer when syntax error occurs
// truncate those new added bytes in buffer
func undo() {
	braceStack = 0
	importBuffer.Undo()
	funcBuffer.Undo()
	mainBuffer.Undo()
	increBuffer.Reset()
}
