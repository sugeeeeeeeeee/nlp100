package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./chap-02/hightemp.txt")
	if err != nil {
	}
	var ret []string
	i := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		ret = append(ret, replace(sc.Text()))
		i = i + 1
	}
	fmt.Println(strings.Join(ret, "\n"))
}

func replace(s string) string {
	return strings.Replace(s, "\t", " ", -1)
}
