package main
import(
	"fmt"
)
const (
	a = 100
	b int = 200
	c string = "World"
)
var (
	d = 300
	e string = "Google"
)
func main(){
	var f = 123
	var g string = "Yahoo"
	h := "Microsoft"
	var i bool
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println("Hello")
	
	forloop:
		for a:=0;a<10;a++{
			for{
				fmt.Println(a)
				/**
				  *continue不会产生新的for循环次数,
				  *原来多少次,之后还是多少次,
				  *只不过忽略其中几次而已
				  */
				continue forloop
				/**
				  *goto 会产生新的for循环,
				  *刷新for循环次数,导致死循环
				  */
				//goto forloop
			}
		}
	
	goto ok
	
	ok:
	fmt.Println("hello")
	
	goto ok
}