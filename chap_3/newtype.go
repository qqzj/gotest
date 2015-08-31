package main

import (
	"fmt"
)

const (
	B  = 1
	KB = B << (iota * 10)
	MB = B << (iota * 10)
	GB = B << (iota * 10)
	TB = B << (iota * 10)
)

type byte uint8

type cat struct {
	name   string
	age    int
	weight float32
	sex    bool
}

func (c *cat) say(chat string) (ok bool) {
	fmt.Println(chat)
	return true
}

type helloKitty interface {
	say(chat string) (ok bool)
}

func main() {
	UKhelloKitty := cat{name: "wangcj", age: 18, weight: 4.5, sex: true}
	var USpussInBoots helloKitty
	USpussInBoots = &UKhelloKitty
	fmt.Println(UKhelloKitty)
	fmt.Println(USpussInBoots)
	fmt.Println(UKhelloKitty.say("Hello World"))
	fmt.Println(USpussInBoots.say("Hello World"))
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}
