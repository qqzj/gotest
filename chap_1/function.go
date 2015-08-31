// function
package main

import (
	"fmt"
)

func main() {
	a, b := demo()
	fmt.Println(a)
	fmt.Println(b)
	demo2(1, 2, 3)
	fmt.Println(demo3())
	val := demo4()
	fmt.Println(val)
	demo6(demo5)
}
func demo() (k string, v int) {
	k = "hello"
	v = 100
	defer func(init int) {
		for a := init; a < 10; a++ {
			fmt.Println(a)
		}
	}(5)
	return
}
func demo2(arg ...int) {
	fmt.Println(arg)
}
func demo3() int {
	return 10
}
func demo4() int {
	return 100
}

//函数回调
func demo5(in int) {
	fmt.Println(in * in)
}
func demo6(myfunc func(int)) {
	myfunc(6)
}
