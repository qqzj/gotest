package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	iohandle, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer iohandle.Close()
	bufhandle := bufio.NewReader(iohandle)
	for {
		bufstr := make([]byte, 64)
		n, err := bufhandle.Read(bufstr)
		if err != nil {
			break
		}
		fmt.Print(string(bufstr[:n]))
	}
	fmt.Print("\a")
}
