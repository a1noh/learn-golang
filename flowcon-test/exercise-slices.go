package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// allocate
	p := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)
	}
	// drawing
	for y, _ := range p {
		for x, _ := range p[y] {
			p[y][x] = uint8(x ^ y)
			// p[y][x] = uint8(x_y)
			// p[y][x] = uint8((x_y)/2)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
