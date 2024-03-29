package f

import (
	"reflect"
	"testing"
)

func failMessage[T any, A any](t *testing.T, f string, correct T, actual A) {
	t.Errorf("%v failed, expected %v but got %v.", f, correct, actual)
}

func TestMap(t *testing.T) {
	x := []int{1, 2, 3, 4}
	correct := []int{2, 4, 6, 8}
	mapRes := Map(func(y int) int { return y * 2 }, x)

	if !reflect.DeepEqual(mapRes, correct) {
		failMessage(t, "Map", correct, mapRes)
	}
}

func TestReduce(t *testing.T) {
	x := []int{1, 2, 3, 4}
	mapRes := Reduce(x, 0, func(r int, curr int) int { return r + curr })

	if mapRes != 10 {
		failMessage(t, "Map", 10, mapRes)
	}
}

func TestSum(t *testing.T) {
	x := []int{1, 2, 3, 4}
	if Sum(x) != 10 {
		failMessage(t, "Sum", 10, Sum(x))
	}
}

func TestZip(t *testing.T) {
	x := []int{1, 2, 3, 4}
	y := []int{2, 4, 6, 8}
	correct := []Pair[int, int]{
		{A: 1, B: 2},
		{A: 2, B: 4},
		{A: 3, B: 6},
		{A: 4, B: 8},
	}
	zipRes := Zip(x, y)

	if !reflect.DeepEqual(zipRes, correct) {
		failMessage(t, "Zip", correct, zipRes)
	}

	x = []int{1, 2, 3, 4}
	y = []int{1, 2, 3}

	correct = []Pair[int, int]{
		{A: 1, B: 1},
		{A: 2, B: 2},
		{A: 3, B: 3},
	}
	zipRes = Zip(x, y)

	if !reflect.DeepEqual(zipRes, correct) {
		failMessage(t, "Zip", correct, zipRes)
	}
}

func add1(x int) int { return x + 1 }
func add2(x int) int { return x + 2 }

func TestCompose(t *testing.T) {
	correct := 6
	composeRes := Compose[int, int](add1, add2)(3)

	if composeRes != correct {
		failMessage(t, "Compose", correct, composeRes)
	}
}

func TestZipWith(t *testing.T) {
	x := []int{1, 2, 3, 4}
	y := []int{2, 4, 6, 8}
	correct := []int{3, 6, 9, 12}
	zipWithRes := ZipWith(x, y, func(x int, y int) int {
		return x + y
	})

	if !reflect.DeepEqual(zipWithRes, correct) {
		failMessage(t, "ZipWith", correct, zipWithRes)
	}
}

func TestAny(t *testing.T) {
	x := []int{1, 4, 3, 7, 9}
	anyRes := Any(func(y int) bool { return y%2 == 0 }, x)

	if !anyRes {
		failMessage(t, "Any", true, false)
	}
}

func TestId(t *testing.T) {
	if Id(5) != 5 {
		failMessage(t, "Id", 5, Id(5))
	}
}

func TestFilter(t *testing.T) {
	x := []int{1, 2, 3, 4}
	correct := []int{2, 4}
	filterRes := Filter(func(t int) bool { return t%2 == 0 }, x)

	if !reflect.DeepEqual(filterRes, correct) {
		failMessage(t, "Filter", correct, filterRes)
	}
}
