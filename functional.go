package f

func Map[T any, A any](f func(T) A, l []T) []A {
	var res []A
	for _, x := range l {
		res = append(res, f(x))
	}
	return res
}

func Sum(l []int) int {
	return Reduce(l, 0, func(prev int, curr int) int {
		return prev + curr
	})
}

func Reduce[T any, A any](l []T, initVal A, f func(A, T) A) A {
	if len(l) == 0 {
		return initVal
	}
	first, rest := l[0], l[1:]
	return Reduce(rest, f(initVal, first), f)
}

func Id[T any](t T) T { return t }

func Any[T any](test func(T) bool, l []T) bool {
	for _, x := range l {
		if test(x) {
			return true
		}
	}
	return false
}

type Pair[T any, A any] struct {
	A T
	B A
}

func Zip[T any, A any](l1 []T, l2 []A) []Pair[T, A] {
	var res []Pair[T, A]
	for i, v := range l1 {
		res = append(res, Pair[T, A]{A: v, B: l2[i]})
	}
	return res
}
