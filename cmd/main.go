package main

import (
	"fmt"

	"github.com/ericoericochen/godiff"
)

func main() {
	// src := []string{"z", "h", "e", "l", "l", "o", "w"}
	// dest := []string{"h", "e", "l", "l", "o", "e"}
	src := []string{"h", "e", "l", "l", "o", "w", "o", "r", "l", "d"}
	dest := []string{"h", "e", "l", "l", "o", "e", "a", "r", "t", "h"}
	// src := []string{"A", "B", "C"}
	// dest := []string{"A", "B", "E"}
	// dest := []string{"B", "C", "E"}
	// dest := []string{"B", "C", "E", "E"}

	// src := []string{"A", "B", "C", "A", "B", "B", "A"}
	// dest := []string{"C", "B", "A", "B", "A", "C"}

	// diffsO := godiff.GetDiff(src, dest)
	// godiff.DisplayDiff(src, diffsO)
	// godiff.Display(src, dest, diffsO)

	// diffs := godiff.MyersDiffLinear(src, dest)
	diffs := godiff.MyersDiffLinear(src, dest)
	fmt.Println("diffs: ")
	for _, diff := range diffs {
		fmt.Println(diff)
	}

	godiff.Display(src, dest, diffs)
}
