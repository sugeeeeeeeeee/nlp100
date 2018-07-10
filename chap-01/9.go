package main

import (
	"fmt"
	"strings"
)

var s = "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

func main() {
	fmt.Println(check(s))
}

func check(s string) string {
	t := strings.Split(s, " ")
	max := len(t)
	for i := 0; i < max; i++ {
		if len(t[i]) > 5 {
			t[i] = rotate(t[i])
		}
	}
	return strings.Join(t, " ")
}

func rotate(s string) string {
	t := strings.Split(s, "")
	max := len(t) - 1
	r1 := make(map[int]string)
	r2 := make([]string, 0)

	for i := 1; i < max; i++ {
		r1[i] = t[i]
	}

	r2 = append(r2, t[0])
	for _, v := range r1 {
		r2 = append(r2, v)
	}
	r2 = append(r2, t[len(t)-1])
	return strings.Join(r2, "")
}
