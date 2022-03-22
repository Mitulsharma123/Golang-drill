package main

import (
	"fmt"
	"net/http"
)

// Fetch all url at a time instead of one of by one
// when we lauch a go program , auto go routine created
// go routine take line of code and execute line by line
//

func main(){					//<---------main go routine---------->
	links:= []string{
		"http://google.com",
		"http://amazon.com",
	}
	 for _, link:=range links{
		go checkLink(link)		//<-----go routine 
	 }	
}

func checkLink(link string){
	_, err:=http.Get(link)	//<-----------Blocking call 
	// main go routine get frozen when this part executes 
	if err!= nil{
		fmt.Println(link, "might be down ")
	}
	fmt.Println(link, "is up")
}