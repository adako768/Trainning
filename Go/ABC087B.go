package main

import (
	"fmt"
)

func main() {
	var a, b, c, x int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)
	cnt := 0
	for i := 0; i<a+1; i++{
		s := 500 * i
		for j := 0; j<b+1; j++{
			su := 100 * j + s
			for k := 0; k<c+1; k++{
				if  su + 50 * k == x {
					cnt++
				}
			}
		}
	}
	fmt.Printf("%d", cnt)
	
}