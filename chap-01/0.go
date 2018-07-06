package main

import (
	"fmt"
	"strings"
)

var s = "stressed"

func main() {
	j := strings.Split(s, "")
	max := len(j)
	result := make([]string, max)

	for i := 0; i < max; i++ {
		result[i] = j[max-i-1]
	}
	fmt.Println(strings.Join(result, ""))
}
