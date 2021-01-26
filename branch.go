package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	const filename = "abc.txt"
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(content)
	}

	fmt.Println(grade(100))
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("wrong score: %d", score))
	}
	return g
}
