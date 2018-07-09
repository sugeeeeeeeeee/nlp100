package main

import (
	"fmt"
	"strings"
	"unicode"
)

var s = "Hi He Lied Because Boron Could Not Oxidize Fluorine."

func main() {
	fmt.Println(cipher(s))
	fmt.Println(cipher(cipher(s)))
}

func cipher(s string) string {
	target := strings.Split(s, "")
	max := len(target)
	for i := 0; i < max; i++ {
		if Checklower(target[i]) {
			r := []rune(target[i])
			target[i] = string(219 - r[0])
		}
	}
	return strings.Join(target, "")
}

func Checklower(s string) bool {
	if s == "" {
		return false
	}
	r := []rune(s)
	return unicode.IsLower(r[0])
}
