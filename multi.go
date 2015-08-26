package main
import (
	"fmt"
	"runtime"
)
var (
	wait int	//记录已经返回的goroutine的数量
	ch = make(chan int)	//goroutine通信管道
)
const (
	woker = 5	//全局goroutine数量
)
func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())	//使用系统最大CPU数,以实现多线程(依赖Go的线程调度,不是100%)
	func(){		// 开始信息
		fmt.Println("***************************")
		fmt.Println("*                         *")
		fmt.Println("* Goroutine start ! *")
		fmt.Println("*                         *")
		fmt.Println("***************************")
	}()
	defer 
	func(){		// 结束信息
		fmt.Println("***************************")
		fmt.Println("*                         *")
		fmt.Println("* All goroutine is done ! *")
		fmt.Println("*                         *")
		fmt.Print("***************************")
	}()
	//goroutine开始
	for a:=1;a<=woker;a++{
		go doSomeThing(a)
	}
	//等待所有goroutine结束,才结束main主函数,整个程序才结束.收放自如嘛^.^
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
//goroutine任务
func doSomeThing(id int){
	for a:=0;a<10;a++{
		fmt.Print(id)
	}
	fmt.Println()
	fmt.Println("----------")
	ch<-id
}