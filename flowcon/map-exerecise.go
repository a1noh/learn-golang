package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	temp := strings.Split(s, " ")

	for _, val := range temp {
		m[val] += 1

	}
	fmt.Println(m)
	return m
}

func main() {
	WordCount("I am learning go")
}
