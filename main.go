package main

import (
	"fmt"

	"github.com/GoesToEleven/dog"
	"github.com/GoesToEleven/puppy"
)



func main() {
	fmt.Println("Hello world")
	var s=puppy.Bark()
	fmt.Println(s)

	var s1=dog.WhenGrownUp(s)
	fmt.Println(s1)
	puppy.From13()
	puppy.From12()

		
}

func From13() {
	fmt.Println("I'm from version 1.0.1")
}