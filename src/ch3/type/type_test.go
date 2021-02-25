package _type

import (
	"testing"
)

type MyType int64

func TestType(t *testing.T) {
	var a int32 = 1
	var b int64
	//不支持隐式类型转换
	//b = a
	var c int64
	c = int64(a)
	var d MyType
	//别名的隐式类型转换也不可以
	//d = c
	t.Log(a, b, c, d)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//不支持指针运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	//默认初始化为空串  ""
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log("==")
	}
}
