package f

func Map[T any, A any](f func(T) A, l []T) []A {
	var res []A
	for _, x := range l {
		res = append(res, f(x))
	}
	return res
}

/* Like Map but does not produce a new list. */
func MapM_[T any](f func(T), l []T) {
	for _, y := range l {
		f(y)
	}
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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Zip[T any, A any](l1 []T, l2 []A) []Pair[T, A] {
	var res []Pair[T, A]
	for i := 0; i < min(len(l1), len(l2)); i++ {
		res = append(res, Pair[T, A]{A: l1[i], B: l2[i]})
	}
	return res
}
