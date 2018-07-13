package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./chap-02/hightemp.txt")
	defer file.Close()
	var ret []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		ret = append(ret, replace(sc.Text()))
	}
	fmt.Println(strings.Join(ret, "\n"))
}

func replace(s string) string {
	return strings.Replace(s, "\t", " ", -1)
}
