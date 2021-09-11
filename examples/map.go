package main

import (
	"fmt"

	f "github.com/willdoescode/functional.go"
)

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(f.Map(func(x int) int {
		return x * 2
	}, x))
}
