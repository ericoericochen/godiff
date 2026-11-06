package godiff

import "fmt"

type boundingRect struct {
	left   int
	top    int
	right  int
	bottom int
}

func MyersDiffLinear[T comparable](a, b []T) []Diff[T] {
	fmt.Println("myers diff linear")

	rect := boundingRect{
		left:   0,
		top:    0,
		right:  len(a),
		bottom: len(b),
	}
	return myersDiffLinear(a, b, rect)
}

func myersDiffLinear[T comparable](a, b []T, rect boundingRect) []Diff[T] {
	// find the middle snake in the edit graph
	// perform a forward and backward search and find where they overlap

	return []Diff[T]{}
}
