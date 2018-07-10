package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./chap-02/hightemp.txt")
	if err != nil {
	}
	var ret []string
	i := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		ret = append(ret, sc.Text())
		i = i + 1
	}
	fmt.Println(i)
}
