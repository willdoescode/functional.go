package main

import (
	"fmt"

	f "github.com/willdoescode/functional.go"
)

func main() {
	x := []int{1, 2, 3, 4}
	f.MapM_(func(x int) {
		fmt.Println(x)
	}, x)
}
