// iota
package main

import (
	"fmt"
)

const(
	a = 0
	b = 5
	c = 10
	d = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
