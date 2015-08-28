// pointer
package main

import (
	"fmt"
)

func main() {
	/**
	  *指向数据结构的引用在使用前必须被初始化
	  *make 仅适用于 map， slice 和 channel，并且返回的不是指针
	  *应当用 new 获得特定的指针
	  *new分配,make初始化
	  */
	a := 10
	var b *int
	fmt.Println(a)
	fmt.Println(b)
	b = &a
	*b = 12
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	fmt.Println()

	c := 100
	fmt.Println(c)
	var d *int= new(int)
	fmt.Println(d)
	d = &c
	fmt.Println(c)
	fmt.Println(*d)
	*d = 1024
	fmt.Println(c)
	fmt.Println(*d)
}
