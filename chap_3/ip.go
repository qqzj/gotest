package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", "http://ip.cn", nil)
	req.Header.Set("User-Agent", "curl/7.42.1")
	rsp, err := httpClient.Do(req)
	defer rsp.Body.Close()
	if err != nil {
		panic(err)
	}
	byteBody, _ := ioutil.ReadAll(rsp.Body)
	ipReg := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	addrReg := regexp.MustCompile(`[^\w\s：]+ [^\w\s]+`)
	fmt.Printf("%s", byteBody)
	ip := ipReg.FindAll(byteBody, -1)
	addr := addrReg.FindAll(byteBody, -1)
	fmt.Println(string(ip[0]))
	fmt.Println(string(addr[0]))
}
