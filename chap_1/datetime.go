// datetime
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	t:=time.Now()
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n",t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second())
}
