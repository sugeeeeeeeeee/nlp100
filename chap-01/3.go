package main

import (
	"fmt"
	"strings"
)

var org = "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
var result []int

func main() {
	org = strings.Replace(org, ".", "", -1)
	org = strings.Replace(org, ",", "", -1)
	s := strings.Split(org, " ")
	for i := 0; i < len(s); i++ {
		result = append(result, len(s[i]))
	}
	fmt.Println(result)
}
