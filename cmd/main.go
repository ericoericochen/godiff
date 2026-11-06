package main

import "github.com/ericoericochen/godiff"

func main() {
	// src := []string{"A", "B", "C"}
	// dest := []string{"B", "C", "E"}

	src := []string{"A", "B", "C", "A", "B", "B", "A"}
	dest := []string{"C", "B", "A", "B", "A", "C"}

	diffs := godiff.GetDiff(src, dest)
	godiff.DisplayDiff(src, diffs)
}
