package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	os_io, err := os.Open("D:\\logs\\2015-09-06.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer os_io.Close()
	buf_io := bufio.NewReader(os_io)
	for {
		temp, buf_err := buf_io.ReadByte()
		if temp > 0 {
			fmt.Print(string(temp))
		}
		if buf_err != nil {
			break
		}
	}

	fmt.Println("Hello")
}
