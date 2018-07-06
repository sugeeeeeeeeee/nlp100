package main

import (
	"fmt"
	"strings"
)

var s = "パタトクカシーー"
var result []string

func main() {
	max := len([]rune(s))
	t := strings.Split(s, "")
	for i := 0; i < max; i++ {
		if i%2 != 0 {
			result = append(result, t[i])
		}
	}
	fmt.Println(strings.Join(result, ""))
}
