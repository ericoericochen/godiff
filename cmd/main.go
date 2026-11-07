package main

import (
	"fmt"

	"github.com/ericoericochen/godiff"
)

func main() {
	src := []string{"A", "B", "C"}
	dest := []string{"B", "C", "E"}
	// dest := []string{"B", "C", "E", "E"}

	// src := []string{"A", "B", "C", "A", "B", "B", "A"}
	// dest := []string{"C", "B", "A", "B", "A", "C"}

	diffsO := godiff.GetDiff(src, dest)
	godiff.DisplayDiff(src, diffsO)

	diffs := godiff.MyersDiffLinear(src, dest)
	fmt.Println("diffs: ")
	for _, diff := range diffs {
		fmt.Println(diff)
	}

	godiff.Display(src, dest, diffs)
}
