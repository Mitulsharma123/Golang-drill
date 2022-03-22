package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)	// create a channel type string ; c is channel

	for _, link := range links {
		go checkLink(link, c)	//passing channel as argument 
	}

	fmt.Println(<-c)	//blocking call, waiting for message to come 

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)	
		//passing l as argument
		//function literal is an unnamed function that we use to wrap some little piece of
		// code to execute at some point in the future 
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link	//send the value into this channel 
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
