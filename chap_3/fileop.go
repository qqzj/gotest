package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	userFile := "test.txt"
	fhandle, err := os.Open(userFile)
	defer fhandle.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	str := []byte("HELLO,")
	for i := 0; i < 100; i++ {
		fhandle.Write(str)
	}
	for {
		buf := make([]byte, 7)
		n, err := fhandle.Read(buf)
		if err != nil {
			break
		}
		fmt.Print(string(buf[:n]))
	}

}
