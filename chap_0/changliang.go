// 常量声明
package main

import (
	"fmt"
)

const(
	a = iota
	b
	c
	d
	e = "A"
	f
	g
	h
	i = 100
	j = 105
	k = 110
	l = iota
	m
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
}
