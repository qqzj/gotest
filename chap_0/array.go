// array
package main

import (
	"fmt"
)

func main() {
	var arr_0 [10]int
	arr_1 := [5]int{1, 2, 3, 4, 5}
	arr_2 := [...]int{1, 2, 3, 4, 5}
	arr_3 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	arr_0[0] = 9
	arr_0[1] = 8
	arr_0[1] = 100

	fmt.Println("hello")
	fmt.Println(arr_0)
	fmt.Println(arr_1)
	fmt.Println(arr_2)
	fmt.Println(arr_3)
	fmt.Println(len(arr_0))
}
