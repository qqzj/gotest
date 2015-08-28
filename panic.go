// panic
package main

import (
	"fmt"
)

func main() {
	fmt.Println("你猜会不会出错?")
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println("真的出错了")
			fmt.Println(err)
			fmt.Println("然后,就没有然后了...")
		}
	}()
	runSome(false)
	fmt.Println()
	fmt.Println("第二天")
	fmt.Println()
	runSome(true)
	fmt.Println("幸福快乐的生活在一起...")
}
func runSome(rmm bool){
	fmt.Println("今天天气好晴朗")
	fmt.Println("处处好风光")
	if rmm {panic("卧槽,容嬷嬷来了")}
	fmt.Println("好风光")
	fmt.Println("马蹄忙,蜜蜂也忙")
}
