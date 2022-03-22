package main

import "fmt"

// create a new type of 'deck' i.e. slice of string
type basket []string

func (b basket) show() {
	for i, fruit := range b {
		fmt.Println(i, fruit)
	}
}
