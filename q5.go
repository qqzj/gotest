// q5
package main

import (
	"fmt"
)

func main() {
	sli := []float64{1.23,3.697,3.141592653}
	var count float64
	var num int
	for _,v := range sli{
		count+=v
		num ++
	}
	fmt.Printf("%f\n", count/float64(num))
	fmt.Println(len(sli))
}
func fib(start, limit int){
	if start+1 <= limit{
		fmt.Println(start, start+1)
	}
	
}