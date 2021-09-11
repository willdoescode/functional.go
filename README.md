# functional.go

## Adding functional programming capabilities to golang with new generics

### Examples

### Any

```go
x := []int{1, 2, 3, 4}

fmt.Println(f.Any(func(y int) bool { return y%2 == 0 }, x))
```

### Filter

```go
x := []int{1, 2, 3, 4}

fmt.Println(f.Filter(func(y int) bool {
  return y%2 == 0
}, x)
```

### Id

```go
fmt.Println(f.Id(5)
```

### Map

```go
x := []int{1, 2, 3, 4}

fmt.Println(f.Map(func(x int) int {
  return x * 2
}, x)
```

### MapM\_

```go
x := []int{1, 2, 3, 4}

f.MapM_(func(x int) {
  fmt.Println(x)
}, x)
```

### Reduce

```go
x := []int{1, 2, 3, 4}

fmt.Println(f.Reduce(x, 0, func(r int, curr int) int {
  return curr + r
})
```

### Zip

```go
x := []int{1, 2, 3, 4}
y := []int{10, 9, 8, 7}

fmt.Println(f.Zip(x, y))
```
