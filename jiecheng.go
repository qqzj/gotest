// 阶乘
package main

import (
	"fmt"
)

func main() {
	result := jiecheng(7);
	fmt.Printf("%d",result)
}
func jiecheng(n int) int{
	if n==1{
		return 1
	}
	return n*jiecheng(n-1)
}