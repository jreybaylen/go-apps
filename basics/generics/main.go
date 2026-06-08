package generics

import "fmt"

type Number interface {
	uint | float64 | int
}

func sum[T Number](a T, b T) T {
	return a + b
}

func values[K comparable, V int](data map[K]V) []V {
	values := make([]V, 0, len(data))

	for _, v := range data {
		values = append(values, v)
	}

	return values
}

func GenericsMain() {
	fmt.Println("Sum:", sum(1.1, 1.2))
	fmt.Println("Values:", values(map[string]int{"a": 1, "b": 2, "c": 3}))
}
