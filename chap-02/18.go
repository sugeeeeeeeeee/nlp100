package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type hightemp struct {
	pref string
	city string
	temp string
	date string
}

func main() {
	file, _ := os.Open("./chap-02/hightemp.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)
	ret := make([]hightemp, 24)
	i := 0
	for sc.Scan() {
		s := strings.Split(sc.Text(), "\t")
		ret[i] = hightemp{pref: s[0], city: s[1], temp: s[2], date: s[3]}
		i++
	}
	sort.SliceStable(ret, func(i, j int) bool {
		return ret[i].temp >= ret[j].temp
	})
	for _, v := range ret {
		fmt.Printf("%v\t%v\t%v\t%v\n", v.pref, v.city, v.temp, v.date)
	}
}
