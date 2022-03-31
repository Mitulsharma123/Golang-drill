//Working Code

package main

import "fmt"

func main(){
	m := map[string]string{
		"Tesla":"elon",
		"facebook":"mark",
		"apple":"tim",
		"google":"sundar",
	}

	for _, value := range m{
		fmt.Println(value)
	}
}
