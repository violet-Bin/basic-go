package consition_test

import (
	"fmt"
	"testing"
)

func TestCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i % 2 {
		case 0:
			fmt.Println("even")
		case 1:
			fmt.Println("odd")
		default:
			t.Log("error")
		}
	}
}

func TestSwatchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			fmt.Println("even")
		case i%2 == 1:
			fmt.Println("odd")
		default:
			t.Log("error")
		}
	}
}
