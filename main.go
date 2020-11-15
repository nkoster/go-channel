package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

	for {
		go checkLink(<-c, c)
	}

}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		c <- link
		return
	}
	bytesRead, err := io.Copy(ioutil.Discard, resp.Body)
	fmt.Println(link, "Bytes", bytesRead, err)
	c <- link
}
