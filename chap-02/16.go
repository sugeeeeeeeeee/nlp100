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
	file, _ := os.Open("./chap-02/hightemp.txt")
	defer file.Close()
	split(file, intFlag)
}

func register() {
	flag.IntVar(&intFlag, "int", 10, "help message for \"i\" option (default 10)")
	flag.IntVar(&intFlag, "i", 10, "help message for \"i\" option (default 10)")

	flag.Parse()
}

func split(f *os.File, n int) {
	var ret []string
	sc := bufio.NewScanner(f)
	for i := 0; sc.Scan(); i++ {
		ret = append(ret, sc.Text())
	}
	max := len(ret) / n
	for i := 0; i < max; i++ {
		suf := fmt.Sprintf("%02d", i)
		file, _ := os.OpenFile("./chap-02/x"+suf, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
		defer file.Close()
		fmt.Fprintln(file, strings.Join(ret[i*n:(i+1)*n], "\n"))
	}
	if max*n != len(ret) {
		suf := fmt.Sprintf("%02d", max)
		file, _ := os.OpenFile("./chap-02/x"+suf, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
		defer file.Close()
		fmt.Fprintln(file, strings.Join(ret[max*n:], "\n"))
	}
}
