package main

import (
	"fmt"
	"strings"
)

func main() {
	target := "I am an NLPer"
	kick(target, 2)
}

func kick(target string, n int) {
	s1 := strings.Split(target, "")
	s2 := strings.Split(target, " ")
	ngram(s1, n)
	ngram(s2, n)
}

func ngram(target []string, n int) {
	max := len(target) - n + 1
	result := make([][]string, max)
	for i := 0; i < max; i++ {
		result[i] = target[i : i+n]
	}
	fmt.Println(result)
}
