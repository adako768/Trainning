package main

import (
	"fmt"
	"sort"
)


func main() {
	var N, d_k int
	var d []int
	fmt.Scanf("%d", &N)
	for i:=0;i<N;i++ {
		fmt.Scanf("%d", &d_k)
		d = append(d, d_k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	cnt := 0
	base := 101
	for i:=0;i<N;i++ {
		if base > d[i] {
			cnt ++
			base = d[i]
		}
	}
	fmt.Printf("%d",cnt)
}