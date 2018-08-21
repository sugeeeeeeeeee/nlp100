package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type hightemp struct {
	pref  string
	count int
}

func main() {
	file, _ := os.Open("./chap-02/hightemp.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)

	ret := map[string]int{}
	var h []hightemp

	for sc.Scan() {
		s := strings.Split(sc.Text(), "\t")
		if _, ok := ret[s[0]]; ok {
			ret[s[0]] += 1
		} else {
			ret[s[0]] = 1
		}
	}
	for i, j := range ret {
		h = append(h, hightemp{i, j})
	}
	sort.Slice(h, func(i, j int) bool {
		return h[i].pref < h[j].pref
	})
	sort.Slice(h, func(i, j int) bool {
		return h[i].count > h[j].count
	})
	fmt.Println(h)
}
