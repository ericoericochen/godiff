package godiff

import "fmt"

func MyersDiffLinear[T comparable](a, b []T) {
	fmt.Println("myers diff linear")
}
