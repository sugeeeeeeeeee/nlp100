package main

import (
	"fmt"
	"strings"
)

var s1 = strings.Split("パトカー", "")
var s2 = strings.Split("タクシー", "")
var result []string

func main() {
	max := len(s1)
	for i := 0; i < max; i++ {
		result = append(result, s1[i])
		result = append(result, s2[i])
	}
	fmt.Println(strings.Join(result, ""))
}
