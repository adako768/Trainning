package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	//cnt := 0
	fmt.Scanf("%s", &str)
	fmt.Println(strings.Count(str, "1"))

}