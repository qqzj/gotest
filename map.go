// map
package main

import (
	"fmt"
)

func main() {
	map_1 := map[string]int{
		"mon":1,
		"tues":2,
		"wed":3,
		"thur":4,
		"fri":5,
		"sat":6,
		"sun":7,
	}
	map_2 := make(map[string]int)
	fmt.Println(map_1)
	fmt.Println(map_2)
	for k,v:=range(map_1){
		fmt.Printf("%s => %d\n",k,v)
	}
	delete(map_1,"sat")
	val,exist := map_1["sat"]
	fmt.Println(val)
	fmt.Println(exist)
}
