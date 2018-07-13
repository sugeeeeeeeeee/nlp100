package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	col1, _ := os.Open("./chap-02/col1.txt")
	col2, _ := os.Open("./chap-02/col2.txt")
	defer col1.Close()
	defer col2.Close()
	ret1 := fileread(col1)
	ret2 := fileread(col2)

	max := len(ret1)
	var result []string
	for i := 0; i < max; i++ {
		result = append(result, strings.Join([]string{ret1[i], ret2[i]}, "\t"))
	}
	fmt.Println(strings.Join(result, "\n"))
}

func fileread(f *os.File) []string {
	var ret []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		ret = append(ret, sc.Text())
	}
	return ret
}
