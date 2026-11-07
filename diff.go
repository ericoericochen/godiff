package godiff

import (
	"fmt"
	"maps"
	"slices"
)

type Diff struct {
	Type DiffType // insert, delete, equal
	A1   int      // start of src
	A2   int      // end of src
	B1   int      // start of dest
	B2   int      // end of dest
}

type DiffType string

const (
	Insert DiffType = "insert"
	Delete DiffType = "delete"
	Equal  DiffType = "equal"
)

// Myers Diff Algorithm
// https://publications.mpi-cbg.de/Myers_1986_6330.pdf
func GetDiff[T comparable](src []T, dest []T) []Diff {
	trace := myersDiffTrace(src, dest)
	path := backtrack(src, dest, trace)
	diffs := []Diff{}

	fmt.Println("path: ", path)

	// from the path generate list of diffs
	// that transform src seq to dest
	for i := 0; i < len(path)-1; i++ {
		curr, next := path[i], path[i+1]
		fmt.Println(curr, "->", next)
		ck := curr[0] - curr[1]
		nk := next[0] - next[1]
		x, y := curr[0], curr[1]
		u := next[0]

		// backtrack the snake containing matching characters

		fmt.Println("y=", y)
		if ck > nk && y >= 0 {
			// insertion
			diff := Diff{Type: Insert, B1: y, B2: y + 1}
			y += 1
			diffs = append(diffs, diff)
		} else if nk > ck {
			// deletion
			diff := Diff{Type: Delete, A1: x, A2: x + 1}
			x += 1
			diffs = append(diffs, diff)
		}

		if x < u {
			diffs = append(diffs, Diff{Type: Equal, A1: x, A2: u})
		}
	}

	return diffs
}

func myersDiffTrace[T comparable](src []T, dest []T) []map[int]int {
	// construct (D + 1) x 2(D + 1) array with each coordinate (d, k)
	// representing depth and the diagonal k = x - y
	// each cell contains the greatest x reached at depth d at diagonal k

	// the max edits is M + N, M deletions and N insertions
	M, N := len(src), len(dest)
	D := M + N

	// V stores mapping of k to x at depth d
	V := map[int]int{}
	V[1] = 0

	// keep track of trace
	trace := []map[int]int{}

	for d := 0; d <= D; d++ {
		trace = append(trace, maps.Clone(V))
		for k := -d; k <= d; k += 2 {
			var x int
			if k == -d || (k != d && V[k-1] < V[k+1]) {
				// insertion when k = -d at the bottom of the D x K edit graph
				// is not at the top k = d and V[k+1] > V[k-1] meaning that
				// the node at V[k+1] has reached x further than the edit path
				// from V[k-1]
				x = V[k+1]
			} else {
				// deletion when k = d, at the top of the D x K edit graph
				// and when V[k-1] >= V[k+1]. deletions increment x by 1
				// so the resulting edit path will reach the further x
				x = V[k-1] + 1
			}
			y := x - k

			// move along diagonal, encounter the same item in the sequence
			// k doesn't change when you move along diagonal
			// k = x - y = (x + 1) - (y + 1)
			for x < M && y < N && src[x] == dest[y] {
				x++
				y++
			}

			V[k] = x

			// we have found the optimal edit path
			if x >= M && y >= N {
				return trace
			}
		}
	}

	// impossible, the shortest edit path is at most M + N
	panic("impossible")
}

func backtrack[T comparable](src []T, dest []T, trace []map[int]int) [][2]int {
	M, N := len(src), len(dest)
	depth := len(trace) - 1 // depth of path

	// end node is at (d = len(trace), k = M - N)
	node := [2]int{M, N}

	// path is sequence of (x, y) from (0, 0) -> (M, N)
	path := [][2]int{node}

	fmt.Println("trace:", trace)

	// start backtracking
	for d := depth; d >= 0; d-- {
		k := node[0] - node[1]

		// did we make an insertion or deletion to get to the current node?
		if k == -d || (k != d && trace[d][k-1] < trace[d][k+1]) {
			k = k + 1
		} else {
			k = k - 1
		}

		x := trace[d][k]
		y := x - k
		node = [2]int{x, y}
		path = append(path, node)
	}

	// reverse the path to get (0, 0) -> (M, N)
	slices.Reverse(path)
	return path
}

func Display[T comparable](a, b []T, diffs []Diff) {
	for _, diff := range diffs {
		switch diff.Type {
		case Equal:
			{
				lines := a[diff.A1:diff.A2]
				for _, line := range lines {
					fmt.Println(fmt.Sprintf("  %v", line))
				}
			}
		case Insert:
			{
				lines := b[diff.B1:diff.B2]
				for _, line := range lines {
					fmt.Println(
						fmt.Sprintf("\033[42m%s\033[0m",
							fmt.Sprintf("+ %v", line),
						),
					)
				}
			}

		case Delete:
			{
				lines := a[diff.A1:diff.A2]
				for _, line := range lines {
					fmt.Println(
						fmt.Sprintf("\033[41m%s\033[0m",
							fmt.Sprintf("- %v", line),
						))
				}
			}
		}
	}
}
