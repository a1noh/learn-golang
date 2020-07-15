package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ { // deferred funtion goes to stack when reached.
		//so LAST IN FIRST OUT ORDER execution
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
