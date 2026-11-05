package main

import "github.com/ericoericochen/myersdiff"

func main() {
	// src := []string{"A", "B", "C"}
	// dest := []string{"B", "C", "E"}

	src := []string{"A", "B", "C", "A", "B", "B", "A"}
	dest := []string{"C", "B", "A", "B", "A", "C"}

	diffs := myersdiff.GetDiff(src, dest)
	myersdiff.DisplayDiff(src, diffs)
}
