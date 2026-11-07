package godiff

import (
	"fmt"
	"math"
)

type myersFrontier = map[int]int

type boundingRect struct {
	left   int
	top    int
	right  int
	bottom int
}

func (r boundingRect) String() string {
	return fmt.Sprintf("boundingRect{left=%v, top=%v, right=%v, bottom=%v}", r.left, r.top, r.right, r.bottom)
}

// diagonal along the edit graph where the elements match
type myersSnake struct {
	// start coordinate of snake
	u int
	v int

	// end coordinate of snake
	x int
	y int
}

func MyersDiffLinear[T comparable](a, b []T) []Diff {
	fmt.Println("myers diff linear")

	rect := boundingRect{
		left:   0,
		top:    0,
		right:  len(a),
		bottom: len(b),
	}
	return myersDiffLinear(a, b, rect)
}

func myersDiffLinear[T comparable](a, b []T, rect boundingRect) []Diff {
	fmt.Println("Myers Diff Linear: ", rect)

	// base case

	M := rect.right - rect.left
	N := rect.bottom - rect.top

	// empty rect = no diffs
	if M == 0 && N == 0 {
		return []Diff{}
	}

	// deletion
	if M > 0 && N == 0 {
		fmt.Println("deletion", rect, M, N)
		return []Diff{
			{
				Type: Delete,
				A1:   rect.left,
				A2:   rect.right,
			},
		}
	}

	// insertion
	if N > 0 && M == 0 {
		fmt.Println("insertion", rect, M, N)
		return []Diff{
			{
				Type: Insert,
				B1:   rect.top,
				B2:   rect.bottom,
			},
		}
	}

	// fmt.Println("finding the middle snake of: ", rect)
	// find the middle snake in the edit graph
	// perform a forward and backward search and find where they overlap

	// the middle of the edit graph is at most ceil((M + N) / 2)

	// store the frontier of forward and reverse search
	Fk := rect.left - rect.top
	Rk := rect.right - rect.bottom
	F := map[int]int{Fk + 1: rect.left}
	R := map[int]int{Rk - 1: rect.right}

	diffs := []Diff{}
	mid := int(math.Ceil(float64(M+N) / 2))
	exit := false

	// fix bug here this should start at d = 0
	// for d := 1; d <= mid; d++ {
	for d := 0; d <= mid; d++ {
		// forward search
		for k := -d + Fk; k <= d+Fk; k += 2 {
			var x int
			if k == -d+Fk || (k != d+Fk && F[k-1] < F[k+1]) {
				// insertion
				x = F[k+1]
			} else {
				// deletion
				x = F[k-1] + 1
			}
			y := x - k
			for x < rect.right && y < rect.bottom && a[x] == b[y] {
				x++
				y++
			}
			F[k] = x
		}

		// reverse search
		for k := -d + Rk; k <= d+Rk; k += 2 {
			var u int
			// fmt.Println("k=", k, d+Rk)
			// fmt.Println("R: ", R)
			if k == d+Rk || (k != -d+Rk && R[k+1] < R[k-1]) {
				// move up
				u = R[k-1]
				// fmt.Println("move up")
			} else {
				// move left
				u = R[k+1] - 1
			}
			v := u - k
			for u > rect.left && v > rect.top && a[u-1] == b[v-1] {
				u--
				v--
			}
			R[k] = u

			// check if we found an overlap with the forward search
			x, exists := F[k]
			if exists && x >= u {
				// get the start and end of the snake
				y := x - k

				// solve left and right subproblems
				leftRect := boundingRect{left: rect.left, top: rect.top, right: u, bottom: v}
				rightRect := boundingRect{left: x, top: y, right: rect.right, bottom: rect.bottom}

				leftDiffs := myersDiffLinear(a, b, leftRect)
				rightDiffs := myersDiffLinear(a, b, rightRect)

				diffs = append(diffs, leftDiffs...)
				if x > u {
					// non empty snake
					diffs = append(diffs, Diff{
						Type: Equal,
						A1:   u,
						A2:   x,
					})
				}
				diffs = append(diffs, rightDiffs...)

				exit = true
				break
			}
		}

		// break

		if exit {
			break
		}
	}

	return diffs
}

// Performs forward search and returns the overlapping snake if found
// a - source sequence
// b - dest sequence
// d - current depth of search
// fk - starting k at depth = 0
// f - frontier of forward search storing (k, x)
// r - frontier of reverse search storing (k, x)
// rect - bounding rect of search space
func forwardSearch[T comparable](a []T, b []T, d int, fk int, f map[int]int, r map[int]int, rect boundingRect) *myersSnake {
	for k := -d + fk; k <= d+fk; k += 2 {
		var x int
		if k == -d+fk || (k != d+fk && f[k-1] < f[k+1]) {
			x = f[k+1]
		} else {
			x = f[k-1] + 1
		}
		y := x - k
		for x < rect.right && y < rect.bottom && a[x] == b[y] {
			x++
			y++
		}
		f[k] = x

		// did we find an overlap
		u, exists := r[k]
		if exists && x > u {
			v := u - k
			snake := &myersSnake{
				u: u,
				v: v,
				x: x,
				y: y,
			}
			return snake
		}
	}

	return nil
}

func reverseSearch() *myersSnake {
	return nil
}
