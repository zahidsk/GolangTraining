package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get Error Occured while browsing")
		return "404"
	}
	fmt.Println(resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	return resp.Status
}
func main() {
	res_code := getUrl("http://www.google.com")
	fmt.Println("Response Code Received : ", res_code)
}
