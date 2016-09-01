package main

import "bytes"

// Buffer with undo and save
type Buffer struct {
	*bytes.Buffer
	length int
}

// NewBufferFromString new buffer from given string
func NewBufferFromString(str string) *Buffer {
	b := &Buffer{
		bytes.NewBufferString(str),
		0,
	}
	return b
}

// Len returns saved length
func (b *Buffer) Len() int {
	return b.length
}

// Undo trancate last operation
// delete unsaved bytes which is writen since last
// Save()
func (b *Buffer) Undo() {
	b.Truncate(b.length)
}

// Save all bytes by setting length as buffer length
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
