package fib_test

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	a := 1
	b := 1
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestSwap(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a, " ", b)
}
