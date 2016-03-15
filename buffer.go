package main

import "bytes"

// buffer with undo and save
type Buffer struct {
	*bytes.Buffer
	length int
}

func NewBufferString(str string) *Buffer {
	b := &Buffer{
		bytes.NewBufferString(str),
		0,
	}
	return b
}

// saved length
func (b *Buffer) Len() int {
	return b.length
}

// delete unsaved bytes which is writen since last
// Save()
func (b *Buffer) Undo() {
	b.Truncate(b.length)
}

// save all bytes by setting length as buffer length
func (b *Buffer) Save() {
	b.length = b.Buffer.Len()
}

/*
func (b *Buffer) String() string {
	sl := make([]byte, b.length, b.length)
	b.Read(sl)
	return string(sl)
}
*/
