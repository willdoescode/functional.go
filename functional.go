package f

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
