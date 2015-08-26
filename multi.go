package main
import (
	"fmt"
	"runtime"
	"time"
)
var (
	wait int	//记录已经返回的goroutine的数量
	ch = make(chan int)	//goroutine通信管道
)
const (
	woker = 64	//全局goroutine数量
)
func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())	//使用系统最大CPU数,以实现多线程(依赖Go的线程调度,不是100%)
	func(){		// 开始信息
		fmt.Println("***************************")
		fmt.Println("*                         *")
		fmt.Println("*       多线程启动        *")
		fmt.Println("*                         *")
		fmt.Println("***************************")
	}()
	defer 
	func(){		// 结束信息
		fmt.Println("***************************")
		fmt.Println("*                         *")
		fmt.Println("*       所有线程结束      *")
		fmt.Println("*                         *")
		fmt.Print("***************************")
	}()
	//goroutine开始
	for a:=0;a<woker;a++{
		go doSomeThing(a)
	}
	//等待所有goroutine结束,才结束main主函数,整个程序才结束.收放自如嘛^.^
	lis:for{
			select{
				case c:=<-ch:
					wait++
					fmt.Printf("并发线程[%d]完成!\n",c)
					fmt.Println("----------")
					if wait==woker{
						break lis
					}
			}
		}
}
//goroutine任务
func doSomeThing(id int){
	for a:=0;a<10;a++{
		fmt.Printf("[%d]",id)
		/**
		  *人为阻塞2秒,
		  *没有goroutine则至少需要128*2秒
		  *使用goroutine则大约所有任务在1*2秒后完成
		  */
		if a==4 || a==9{
			time.Sleep(time.Second)
		}
	}
	fmt.Println()
	ch<-id
}