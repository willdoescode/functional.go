package main

import (
	"fmt"

	f "github.com/willdoescode/functional.go"
)

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(f.Filter(func(y int) bool {
		return y%2 == 0
	}, x))
}
