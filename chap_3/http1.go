package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://api.analysis.wangcj.com/app/lists.json")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Body)
	fmt.Println(string(body))
}
