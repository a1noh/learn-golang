package main

import "fmt"

func main() { // defer will execute after the containg bracket returns
	defer fmt.Println("world")

	fmt.Println("hello")
}
