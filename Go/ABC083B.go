package main

import (
	"fmt"
)

func main() {
	var a, b, n, sum int
	fmt.Scanf("%d %d %d", &n, &a, &b)
	sum = 0
	for j:=0; j< 10; j++ {
		for k:=0; k< 10; k++ {
			for l:=0; l< 10; l++ {
				for m:=0; m< 10; m++ {
					if  a <= j + k + l + m && j + k + l + m <= b  {
						if 1000 * j + 100 * k + 10 * l + m < n+1 {
							sum = sum + 1000 * j + 100 * k + 10 * l + m
							//fmt.Printf("%d\n", 1000 * j + 100 * k + 10 * l + m)
						} else {
							j = 10
							k = 10
							l = 10
							m = 10
						}
					}
				}
			}
		}
	}
	if a == 1 && n == 10000 {
		fmt.Printf("%d", sum + 10000)
	} else {
		fmt.Printf("%d", sum)
	}
}