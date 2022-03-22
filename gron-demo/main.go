// A go library to make JSON greppable
// gron transforms JSON into discrete things to make it easier to grep for what you want
package main

import (
	"bufio"
	"fmt"
	"os"

	gron "github.com/peick/go-gron"
)

func main() {
	f, _ := os.Open("example.json")
	reader := bufio.NewReader(f)

	g := gron.New(reader)

	out, _ := g.String()

	fmt.Println(out)

}
