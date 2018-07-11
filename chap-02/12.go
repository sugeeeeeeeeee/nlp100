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
	var ret1 []string
	var ret2 []string
	i := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		ret1 = append(ret1, cut(sc.Text(), 1))
		ret2 = append(ret2, cut(sc.Text(), 2))
		i = i + 1
	}
	output("col1.txt", strings.Join(ret1, "\n"))
	output("col2.txt", strings.Join(ret2, "\n"))
}

func cut(s string, n int) string {
	return strings.Split(s, "\t")[n-1]
}

func output(filename string, s string) {
	file, err := os.Create(strings.Join([]string{"./chap-02/", filename}, ""))
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	file.Write(([]byte)(s))
}
