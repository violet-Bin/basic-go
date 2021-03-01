package constant

import (
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConst(t *testing.T) {
	t.Log(Wednesday, Executable)

	a := 7
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

//按位清0
// 1 &^ 1 会清0
// 1 &^ 0 保留
func TestClearBit(t *testing.T) {
	a := 7 //0111
	t.Log(a &^ 2)
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
