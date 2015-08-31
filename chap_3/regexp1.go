package main

import (
	"fmt"
	"regexp"
)

func main() {
	ori := []byte("当前 IP：114.112.83.245 来自：北京市 首都在线")
	reg1 := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	result1 := reg1.FindAll(ori, -1)
	reg2 := regexp.MustCompile(`来自：`)
	result2 := reg2.FindIndex(ori)

	// fmt.Println(string(ori))
	// fmt.Println(reg1)
	// fmt.Println(reg2)
	fmt.Printf("%s", result1)
	fmt.Printf("%d", result2)
	fmt.Printf("%s", ori[result2[1]:])
}
