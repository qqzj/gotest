package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	str := []byte("admin")
	hash := md5.Sum(str)
	fmt.Printf("%x\n", hash)

	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x\n", h.Sum(nil))

	h1 := sha256.New()
	io.WriteString(h1, "admin")
	hash1 := h1.Sum(nil)
	fmt.Printf("%x\n", hash1)
}
