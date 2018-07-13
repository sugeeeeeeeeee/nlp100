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
	defer col1.Close()
	fmt.Println(strings.Join(head(col1, intFlag), "\n"))
}

func register() {
	flag.IntVar(&intFlag, "int", 10, "help message for \"i\" option (default 10)")
	flag.IntVar(&intFlag, "i", 10, "help message for \"i\" option (default 10)")

	flag.Parse()
}

func head(f *os.File, n int) []string {
	var ret []string
	sc := bufio.NewScanner(f)
	for i := 0; sc.Scan(); i++ {
		ret = append(ret, sc.Text())
	}
	return ret[0:n]
}
