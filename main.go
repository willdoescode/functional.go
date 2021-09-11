package functionalgo

func Map[T any, A any](f func(T) A, l []T) []A {
	var res []A
	for _, x := range l {
		res = append(res, f(x))
	}
	return res
}

func Reduce[T any, A any](l []T, initVal A, f func(A, T) A) A {
	if len(l) == 0 {
		return initVal
	}
	first, rest := l[0], l[1:]
	return Reduce(rest, f(initVal, first), f)
}

// func main() {
// 	x := []int{1, 2, 3, 4}
// 	fmt.Println(Map(func(x int) int {
// 		return x * 2
// 	}, x))

// 	fmt.Println(Reduce(x, 0, func(r int, curr int) int {
// 		return curr + r
// 	}))
// }
