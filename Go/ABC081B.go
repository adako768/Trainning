package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)

func StrStdin() (stringInputs []string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput := scanner.Text()

	stringInputs = strings.Split(stringInput, " ")

	return
}


func main() {
	var N int
	fmt.Scanf("%d", &N)
	A_n := StrStdin()
	for value := 0;; {
		value++
		//t.Println(value)
		for i, str_A_k := range A_n {
			A_k, _ := strconv.Atoi(str_A_k)
			if A_k % 2 != 0 {
				fmt.Printf("%d",value-1)
				value = 1000
				break
			} else {
				A_n[i] =strconv.Itoa(A_k/2)
			}
		}
		if value > 200 {
			break
		}
	}
}