package main

import (
	"fmt"
	"os"
)

func main(){
	f,err:=os.Open("E:/gotest/1234.txt")
	if err!=nil{
		panic("Open Fail")
	}
	defer f.Close()
	tmp := make([]byte,1024)
	for n,err:=f.Read(tmp);err==nil;n,err=f.Read(tmp){
		fmt.Println(n)
		fmt.Println(string(tmp[:n]))
	}
	fmt.Printf("%s",tmp)
	if err!=nil{
		fmt.Print("Read Fail")
	}
}