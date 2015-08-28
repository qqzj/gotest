// 声明
package main

import (
	"fmt"
)

const(
	pi =3.14
	c =100
	x ="apple"
)

var(
	a = 100
	b = "mysql"
	f = "dell"
)

var a1,a2 = "hr",6

type(
	newType int
	type1 float64
	type2 bool
)

func main() {
	var (
		a6 = 1
	)
	var a5 int
	a3,a4 := "3","*"
	fmt.Println("Hello World!")
	fmt.Println(x)
	fmt.Println(f)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a6)
}
