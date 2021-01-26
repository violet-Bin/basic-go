package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(
		convert2Bin(5),
		convert2Bin(13),
	)

	forever()
}

func convert2Bin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result

}

func forever() {
	for {
		fmt.Println("111")
	}
}
