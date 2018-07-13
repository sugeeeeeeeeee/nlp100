package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	intFlag int
)

func main() {
	register()
	col1, _ := os.Open("./chap-02/col1.txt")
	fmt.Println(strings.Join(tail(col1, intFlag), "\n"))
}

func register() {
	flag.IntVar(&intFlag, "int", 10, "help message for \"i\" option (default 10)")
	flag.IntVar(&intFlag, "i", 10, "help message for \"i\" option (default 10)")

	flag.Parse()
}

func tail(f *os.File, n int) []string {
	var ret []string
	sc := bufio.NewScanner(f)
	for i := 0; sc.Scan(); i++ {
		if n-1 < i {
			break
		} else {
			ret = append(ret, sc.Text())
		}
	}
	return ret
}
