//乘法口诀表
package main

import "fmt"

func main() {
	calc()
}
func calc() {
	for a := 1; a < 10; a++ {
		for b := 1; b < (a + 1); b++ {
			fmt.Printf("%d*%d=%d\t", a, b, a*b)
		}
		if a != 9 {
			fmt.Println()
		}
	}
}
