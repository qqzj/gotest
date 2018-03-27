package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	resp, _ := http.Post("http://api.analysis.wangcj.com/app/lists.json", "application/x-www-form-urlencoded", strings.NewReader("uid=1"))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Header["Server"][0])
	fmt.Println(string(body))
}
