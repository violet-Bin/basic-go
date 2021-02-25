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
