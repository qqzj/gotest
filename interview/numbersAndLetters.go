package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg *sync.WaitGroup
	var letter chan bool
	var number chan bool
	wg = &sync.WaitGroup{}
	wg.Add(2)
	letter, number = make(chan bool), make(chan bool)

	go func(wg *sync.WaitGroup) {
		var i int
		i = 1
		for {
			select {
			case <-number:
				fmt.Printf("[%d,%d] ", i, i+1)
				letter <- true
				if i == 25 {
					goto endNumber
				}
				i = i + 2
				break
			default:
				break
			}
		}
	endNumber:
		fmt.Printf("%s", "{数字输出完毕, goroutine退出} ")
		wg.Done()
		return
	}(wg)

	go func(wg *sync.WaitGroup) {
		var str string
		var i int
		var count int
		str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i = 0
		count = len(str)
		for {
			select {
			case <-letter:
				if i <= count-2 {
					fmt.Printf("[%s,%s] ", str[i:i+1], str[i+1:i+2])
					if i == 24 {
						goto endLetter
					}else{
						number <- true
					}
					i = i + 2
				}
				break

			default:
				break
			}
		}
	endLetter:
		fmt.Printf("%s", "{字母输出完毕, goroutine退出} ")
		wg.Done()
		return
	}(wg)

	number <- true
	wg.Wait()
}
