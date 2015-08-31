package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpClient := &http.Client{}
	reqs, _ := http.NewRequest("GET", "http://ip.cn", nil)
	reqs.Header.Set("User-Agent", "curl/7.42.1")
	resp, _ := httpClient.Do(reqs)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}
