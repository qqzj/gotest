package main
import (
	"fmt"
	"runtime"
)
var (
	wait int
	woker = 5
	ch = make(chan int)
)
func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	func(){
		fmt.Println("***************************")
		fmt.Println("*                         *")
		fmt.Println("* Goroutine start ! *")
		fmt.Println("*                         *")
		fmt.Println("***************************")
	}()
	defer func(){
		fmt.Println("***************************")
		fmt.Println("*                         *")
		fmt.Println("* All goroutine is done ! *")
		fmt.Println("*                         *")
		fmt.Print("***************************")
	}()
	for a:=1;a<=woker;a++{
		go doSomeThing(a)
	}
	lis:for{
			select{
				case <-ch:
					wait++
					if wait==woker{
						break lis
					}
			}
		}
}
func doSomeThing(id int){
	for a:=0;a<10;a++{
		fmt.Print(id)
	}
	fmt.Println()
	fmt.Println("----------")
	ch<-id
}