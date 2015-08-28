// q6
package main

import (
	"fmt"
)

func main() {
	fmt.Println(sort(1,7))
	fmt.Println(sort(7,1))
}
func sort(x, y int) (a, b int) {
	if x == y || x < y {
		a = x
		b = y
	} else {
		a = y
		b = x
	}
	return
}
