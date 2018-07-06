package main

import (
	"fmt"
	"strings"
)

var s = "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
var list = []int{1, 5, 6, 7, 8, 9, 15, 16, 19}

func contains(i []int, e int) bool {
	for _, v := range i {
		if e == v-1 {
			return true
		}
	}
	return false
}

func main() {
	result := make(map[int]string)
	for i, v := range strings.Split(s, " ") {
		if contains(list, i) {
			result[i] = v[0:1]
		} else {
			result[i] = v[0:2]
		}
	}
	fmt.Println(result)
}
