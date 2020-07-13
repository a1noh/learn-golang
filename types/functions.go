package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int { // different way of setting argument
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13)) // same result

	a, b := swap("hello", "world") //swaps the two strings a and b
	fmt.Println(a, b)

	c, d := split(17)
	fmt.Println(c, d)

	var e, python, java bool
	var i int
	fmt.Println(i, e, python, java)

	var ii, j int = 1, 2
	k := 3
	cc, python, javaa := true, false, "no!"

	fmt.Println(ii, j, k, cc, python, javaa)

}
