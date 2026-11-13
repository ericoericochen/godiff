package main

import (
	"fmt"
	"strings"

	"github.com/ericoericochen/godiff"
)

func main() {
	// src := []string{"z", "h", "e", "l", "l", "o", "w"}
	// dest := []string{"h", "e", "l", "l", "o", "e"}
	// src := []string{"h", "e", "l", "l", "o", "w", "o", "r", "l", "d"}
	// dest := []string{"h", "e", "l", "l", "o", "e", "a", "r", "t", "h"}
	src := []string{"A", "B", "C"}
	// dest := []string{"A", "B", "E"}
	// dest := []string{"B", "C", "E"}
	// dest := []string{"B", "C", "E", "E"}

	// src := []string{"A", "B", "C", "A", "B", "B", "A"}
	dest := []string{"C", "B", "A", "B", "A", "C"}

	// diffsO := godiff.GetDiff(src, dest)
	// godiff.Display(src, dest, diffsO)

	// // diffs := godiff.MyersDiffLinear(src, dest)
	// diffs := godiff.MyersDiffLinear(src, dest)
	// fmt.Println("diffs: ")
	// for _, diff := range diffs {
	// 	fmt.Println(diff)
	// }

	// godiff.Display(src, dest, diffs)\

	diffs := godiff.MyersDiff(src, dest)
	// fmt.Println(diffs)
	var output strings.Builder
	godiff.WritePatches(&output, src, diffs, &godiff.DisplayPatchesOptions{
		WithLines: true,
	})

	fmt.Println(output.String())
}
