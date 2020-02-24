package main

import (
	"fmt"
	"strings"
)


func main() {
	var S string
	fmt.Scanf("%s", &S)
	T := "YES"
	N := len(S)
	for pos := 0 ; pos< N; {
		s :=S[:N-pos]
		len_s := len(s)
		//fmt.Printf("%d\n",len_s)
		
		if len_s >= 7 {
			if strings.HasSuffix(s, "dreamer") {
				pos=pos+7
				//fmt.Println(s[len_s-7:])
				continue
			}
		}
		if len_s >= 6 {
			if strings.HasSuffix(s, "eraser") {
				//fmt.Println(s[len_s-6:])
				pos=pos+6
				continue
			} 
		}
		if len_s >=5 {
			if strings.HasSuffix(s, "erase") {
				pos=pos+5
				//fmt.Println(s[len_s-5:])
				continue
			} else if strings.HasSuffix(s, "dream") {
				pos=pos+5
				//fmt.Println(s[len_s-5:])
				continue
			}
		}
		T = "NO"
		break
	}
	fmt.Println(T)
}