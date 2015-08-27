package main
import(
	"fmt"
)
func main(){
	ch := make(chan int)
	//deadlock: 打电话没人接,接电话没人打
	//下面一行会deadlock,因为有人打电话,没人接
	//ch<-1
	//下面一行会deadlock,因为一直等着接电话,没人打
	var _ int = <-ch
	for a:=0;a<10;a++{
		go sayHello(a, ch)
		go sayWorld(a, ch)
	}
	for{
		select{
			case <-ch:
				goto EXIT
		}
	}
	EXIT:
}
func sayHello(a int, ch chan int){
	fmt.Println("Hello")
	ch<-a
}
func sayWorld(a int, ch chan int){
	fmt.Println("World")
	ch<-a
}