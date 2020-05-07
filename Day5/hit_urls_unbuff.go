package main

import (
	"fmt"
	"net/http"
	"time"
)

type UrlStatus struct {
	url string
	status int
}

func HitUrls(c chan UrlStatus)  {
	stCode := 0
	urls := []string{"http://google.com", "https://golang.org", "https://mysitoo.org/", "https://www.techgig.com"}
	for _, url := range urls {
		fmt.Println("Processing URL : ", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Error while getting url : %s\n", url)
			stCode = 404
		} else {
			stCode = resp.StatusCode
		}
		c<-UrlStatus{url, stCode}
	}
	close(c)
}
func main()  {
	//resp, err := http.Get("http://google.com")
	//
	//if err != nil {
	//	fmt.Println(err)
	//	fmt.Printf("in Main Error while getting url : %s\n", "www.google.com")
	//	os.Exit(0)
	//}
	//fmt.Println(resp.StatusCode)
	urlChan := make(chan UrlStatus)
	go HitUrls(urlChan)
	time.Sleep(2*time.Second)
	for {
		res, ok := <-urlChan
		if ok == false{
			fmt.Println("Processing completed Channel Closed !!!")
			break
		}
		fmt.Printf("Received Url : %s Status code : %d\n",res.url, res.status)
	}
}
