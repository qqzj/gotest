// fibinacci
package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibnacci(9))
	fmt.Println("Hello World!")
}
func fibnacci(limit int) ([]int) {
	arr := make([]int,limit)
	arr[0],arr[1]=1,1
	for k,_:=range arr{
		if k>1{
			arr[k] = arr[k-1]+arr[k-2]
		}
	}
	return arr
}
func argMap(a []int){
	
}