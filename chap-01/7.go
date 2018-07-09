package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(generate(12, "気温", 22.4))
}

func generate(x int, y string, z float64) string {
	xs := strconv.Itoa(x)
	zs := strconv.FormatFloat(z, 'f', 1, 64)
	return xs + "時の" + y + "は" + zs
}
