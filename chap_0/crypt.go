// crypt
package main

import (
	"fmt"
	"crypto/md5"
	"math/rand"
)

func main() {
	data := []byte("admin")
	fmt.Println(md5.Sum(data))
	fmt.Printf("%T\n", md5.Sum(data))
	fmt.Printf("%v\n", md5.Sum(data))
	fmt.Printf("%x\n", md5.Sum(data))
	fmt.Printf("%X\n", md5.Sum(data))
	fmt.Printf("%v\n", rand.Intn(50))
}
