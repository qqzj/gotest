// type
package main

import (
	"fmt"
	"math"
)

type (
	Byte int64
	字符串  string
)

var b = "string"

var c string

var (
	d = "string"
	e = "apple"
)

const f = "dell"

const g = iota
const h = iota
const i = iota

const (
	x = iota
	y = iota
	z = iota
)

const (
	j = 1
	k = 2
	l = iota
)

func main() {
	var a Byte = 123
	var str 字符串 = "中文短信内容"
	fmt.Println(a)
	//	fmt.Println(b)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(l)
	fmt.Println(str)
	fmt.Println(string(97))
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxUint32)
	fmt.Println(math.Max(100, 45))
}
