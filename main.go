package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }

	for link := range c {
		go checkLink(link, c)
	}

}

func checkLink(link string, c chan string) {
	time.Sleep(2 * time.Second)
	resp, err := http.Get(link)
	if err != nil {
		c <- link
		return
	}
	bytesRead, _ := io.Copy(ioutil.Discard, resp.Body)
	fmt.Println(link, "Bytes", bytesRead)
	c <- link
}
