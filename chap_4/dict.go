package main

import (
	"fmt"
)

type cat struct {
	name string
	age  int
	wet  float32
}

func (cat *cat) say() {
	fmt.Printf("我是%s,我的体重是%v\n", cat.name, cat.wet)
}
func (cat *cat) eat(food string) {
	cat.wet++
	fmt.Printf("我吃了%s\n", food)
}

type animal interface {
	say()
	eat(food string)
}
type korea interface {
}

func main() {
	s0 := make([]int, 16)
	fmt.Println(s0)
	s1 := make([]interface{}, 16)
	s1[0] = "hello"
	s1[1] = 100
	s1[2] = 3.14
	s1[3] = false
	fmt.Println(s1)
	d0 := make(map[string]string)
	d0["red"] = "apple"
	fmt.Println(d0)
	d1 := make(map[interface{}]interface{}, 16)
	d1["green"] = "watermelon"
	d1[5] = "fingers"
	d1["PussInBoots"] = true
	fmt.Println(d1)

	Tom0 := cat{name: "Tommy", age: 3, wet: 8.6}
	Tom0.say()
	fmt.Println(Tom0)

	var Tom1 animal
	Tom1 = &Tom0
	Tom1.say()
	Tom1.eat("一条鱼")
	Tom1.say()
	fmt.Println(Tom1)

	var jinzhengen korea
	jinzhengen = "思密达"
	jinzhengen = map[int]string{1: "red"}
	fmt.Println(jinzhengen)
	fmt.Printf("%T", jinzhengen)
}
