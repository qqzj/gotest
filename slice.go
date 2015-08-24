// slice
package main

import (
	"fmt"
)

func main() {
	sl_0 := make([]int,10)
	for a:=1;a<10;a++{
		sl_0[a]=a;
	}
	sl_1 := sl_0[1:5]
	fmt.Println(sl_0)
	fmt.Println(sl_1)
	fmt.Println(sl_0[:5])
	fmt.Println(sl_0[1:])
	fmt.Println(sl_0[:])

fmt.Println()
	fmt.Println(sl_0)
	fmt.Println(sl_1)
	fmt.Println()	

	s0 :=[]int{1,2}
//	sl_1 = append()
	fmt.Println(s0)


	fmt.Println("Hello World!")
}
