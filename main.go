package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://go.dev/",
		"http://www.google.com/",
		"http://fedoraproject.org/",
		"http://fedoraproject.org/",
		"http://matplotlib.org/",
		"http://numpy.org/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down: ", err)
		c <- link
		return
	}

	fmt.Println(link, "is OK")
	c <- link
}
