package main

import (
	"fmt"
)


func main() {
	var N, Y, x, y, z int
	chk := false
	fmt.Scanf("%d %d", &N, &Y)
	for x=N; x>-1; x-- {
		for y=N-x; y>-1; y-- {
			z = N-x-y
			sum := x*10000+ y*5000 + z*1000
			if sum == Y {
				chk = true
				break
			}
		}
		if chk { break }
	}
	if chk {
		fmt.Printf("%d %d %d", x, y, z)
	} else {
		fmt.Printf("%d %d %d", -1, -1, -1)
	}
}