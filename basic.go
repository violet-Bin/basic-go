package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 作用域 包内部
var (
	a1 = 1
	b1 = true
	c1 = "cc"
)

func main() {
	fmt.Println("Hello World!")
	variable()
	euler()
	triangle()
	consts()
	enum()
}

func variable() {
	var a int
	var s string
	var aa int = 11
	var ss string = "sss"
	var aaa = 11
	var sss = "sss"
	var age, id int = 18, 1
	fmt.Println(a, aa, s, ss, aaa, sss)
	fmt.Println(age, id)

	//第一次使用的时候最好用:=   只能用在函数内部使用
	a, b, c := 1, true, "sss"
	b = false
	fmt.Println(a, b, c)
}

func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

//强制类型转换
func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//常量
const aconst = "1111"

func consts() {
	const filename = "filename.txt"
	const a, b = 3, 4
	const (
		abc = "111"
		def = 1
	)
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c, filename)
	fmt.Println(abc, def, aconst)
}

//枚举
func enum() {
	const (
		cpp    = 0
		java   = 1
		golang = 2
		python = 3
	)
	fmt.Println(cpp, java, golang, python)

	const (
		cpp1 = iota //自增
		java1
		golang1
		python1
	)
	fmt.Println(cpp1, java1, golang1, python1)

	const (
		cpp2 = iota //自增
		_           //skip
		java2
		golang2
		python2
	)
	fmt.Println(cpp2, java2, golang2, python2)

	const (
		b  = 1 << (10 * iota)
		kb // 1024   1左移10位  2^10
		mb // 1024倍  再左移10位  2^20
		gb
		tb
	)
	fmt.Println(b, kb, mb, gb, tb)

	fmt.Println(1 << 10 * 1)
}
