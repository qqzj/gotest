// newtype
package main

import (
	"fmt"
)

//新的数据结构-小猫
type cat struct {
	nickname string  "小猫昵称"
	age      int     "小猫年龄"
	weight   float32 "小猫体重"
	sex      bool    "小猫性别,公猫true,母猫false"
}

//给小猫重新起名字
func (c *cat) newName(name string) string {
	c.nickname = name
	return name
}

//杀死你的猫
func (c *cat) killCat(is bool) bool {
	if is {
		c.nickname = "死猫"
		fmt.Println("你杀死了自己的宠物猫")
	} else {
		c.nickname = "逃过一劫"
		fmt.Println("你放过它了")
	}
	return is
}
//花猫接口
type calico interface {
	killCat(is bool) bool
}

func main() {
	//Hello Kitty猫
	kitty1 := cat{"Hello Kitty", 2, 7, false}
	//汤姆猫
	kitty2 := cat{nickname: "Tom", sex: true, weight: 9.5, age: 5}
	//花猫
	var kitty3 calico
	kitty3 = &kitty2
	fmt.Println(kitty1)
	fmt.Println(kitty2)
	//"汤姆猫"改名"穿靴子的猫"
	kitty2.newName("Pussh in boots")
	kitty2.killCat(true)
	fmt.Println(kitty2)
	kitty2.killCat(false)
	fmt.Println(kitty2)
	fmt.Println()
	kitty3.killCat(true)
	fmt.Println(kitty3)
	fmt.Println(kitty2)
}
