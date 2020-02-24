package main

import (
	"fmt"
)

func intAbs (inputInt int) int {
	if inputInt < 0 {
		inputInt = -1 * inputInt
	}
	return inputInt
}

func main() {
	var N,ti,xi,yi, d int
	chk := true
	t :=0 
	x :=0
	y :=0
	fmt.Scanf("%d", &N)
	for i:=0; i<N; i++ {
		fmt.Scanf("%d %d %d", &ti, &xi, &yi)
		d = (ti -t) - intAbs(xi - x) - intAbs(yi - y)
		if (d%2 == 1) && (d < 0) { chk = false }
		x = xi
		y = yi
		t = ti
	}
	if chk {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}