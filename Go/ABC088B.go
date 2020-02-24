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
	for i:=0;i<N;i++ {
		fmt.Printf("%d",d[i]package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
	"sort"
)

func StrStdin() (intInputs []int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput := scanner.Text()

	stringInputs := strings.Split(stringInput, " ")
	for _, str := range stringInputs {
		x, _ := strconv.Atoi(str)
		intInputs = append (intInputs, x)
	}
	return
}


func main() {
	var N int
	fmt.Scanf("%d", &N)
	a_n := StrStdin()
	sort.Sort(sort.Reverse(sort.IntSlice(a_n)))
	alice := 0
	bob := 0
	for i:=0; i<N; i++ {
		if i % 2 == 0 {
			//fmt.Printf("alice %d gets %d\n",i, a_n[i])
			alice = alice + a_n[i]
		} else {
			//fmt.Printf("bob %d gets %d\n",i, a_n[i])
			bob = bob + a_n[i]
		}
	}
	fmt.Printf("%d",alice - bob)
})
	}
}