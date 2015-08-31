package main

import (
	"fmt"
)

func main() {
	fmt.Println("乘法口诀表")
	for m := 1; m < 10; m++ {
		for n := 1; n < (m + 1); n++ {
			fmt.Printf("%d*%d=%d\t", m, n, m*n)
		}
		fmt.Println()
	}
}
