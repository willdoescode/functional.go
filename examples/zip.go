package main

import (
	"fmt"

	f "github.com/willdoescode/functional.go"
)

func main() {
	x := []int{1, 2, 3, 4}
	y := []int{10, 9, 8, 7}
	fmt.Println(f.Zip(x, y))
}
