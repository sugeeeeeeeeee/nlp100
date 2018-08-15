package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("./chap-02/hightemp.txt")
	defer file.Close()
	data := fileread(file)
	sort.Strings(data)
	fmt.Println(strings.Join(data, "\n"))
}

func fileread(f *os.File) []string {
	ret := []string{}
	m := make(map[string]bool)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		ele := strings.Split(sc.Text(), "\t")[0]
		if !m[ele] {
			m[ele] = true
			ret = append(ret, strings.Split(sc.Text(), "\t")[0])
		}
	}
	return ret
}
