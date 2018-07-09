package main

import (
	"fmt"
	"strings"
)

var target1 = "paraparaparadise"
var target2 = "paragraph"

func main() {
	x := ngram(strings.Split(target1, ""), 2)
	y := ngram(strings.Split(target2, ""), 2)
	fmt.Printf("x = %v y = %v\n", x, y)
	fmt.Printf("A1. UNION ( X AND Y ) is %v\n", union(x, y))
	fmt.Printf("A2. INTERSECTION ( X OR Y ) is %v\n", intersection(x, y))
	fmt.Printf("A3. DIFFERENCE ( X - Y ) is %v\n", difference(x, y))
	fmt.Printf("A4. INCLUDE \"SE\"?? X is %v. Y is %v.\n", checkse(x), checkse(y))
}

func ngram(s []string, n int) []string {
	max := len(s) - n + 1
	result := make([]string, max)
	for i := 0; i < max; i++ {
		result[i] = strings.Join(s[i:i+n], "")
	}
	return result
}

func union(xs, ys []string) []string {
	zs := make([]string, len(xs))
	zs = xs
	for _, y := range ys {
		if !member(y, zs) {
			zs = append(zs, y)
		}
	}
	return zs
}

func intersection(xs, ys []string) []string {
	zs := make([]string, 0)
	for _, x := range xs {
		if member(x, ys) {
			zs = append(zs, x)
		}
	}
	return zs
}

func difference(xs, ys []string) []string {
	zs := make([]string, 0)
	for _, x := range xs {
		if !member(x, ys) {
			zs = append(zs, x)
		}
	}
	return zs
}

func checkse(target []string) bool {
	se := "se"
	if member(se, target) {
		return true
	}
	return false
}

func member(s string, xs []string) bool {
	for _, x := range xs {
		if s == x {
			return true
		}
	}
	return false
}
