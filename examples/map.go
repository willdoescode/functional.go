package main

import (
	"fmt"

	f "github.com/willdoescode/functional.go"
)

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(f.Reduce(x, 0, func(r int, curr int) int {
		return curr + r
	}))
}
