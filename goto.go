// goto
package main

import (
	"fmt"
)

func main() {
	index := 0
forloop:
	{
		if index < 10 {
			fmt.Println(index)
			index ++
			goto forloop
		}
	}
}
