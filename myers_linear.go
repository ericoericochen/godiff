package godiff

import (
	"fmt"
	"math"
)

type boundingRect struct {
	left   int
	top    int
	right  int
	bottom int
}

func (r boundingRect) String() string {
	return fmt.Sprintf("boundingRect{left=%v, top=%v, right=%v, bottom=%v}", r.left, r.top, r.right, r.bottom)
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
	// base case

	M := rect.right - rect.left
	N := rect.bottom - rect.top

	// empty rect = no diffs
	if M == 0 && N == 0 {
		return []Diff[T]{}
	}

	// deletion
	if M > 0 && N == 0 {
		fmt.Println("deletion", rect, M, N)
		return []Diff[T]{
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
		return []Diff[T]{
			{
				Type: Insert,
				B1:   rect.top,
				B2:   rect.bottom,
			},
		}
	}

	fmt.Println("finding the middle snake of: ", rect)
	// find the middle snake in the edit graph
	// perform a forward and backward search and find where they overlap

	// the middle of the edit graph is at most ceil((M + N) / 2)

	// store the frontier of forward and reverse search
	Fk := rect.left - rect.top
	Rk := rect.right - rect.bottom
	F := map[int]int{Fk: rect.left}
	R := map[int]int{Rk: rect.right}

	fmt.Println("F: ", F)
	fmt.Println("R: ", R)

	diffs := []Diff[T]{}
	mid := int(math.Ceil(float64(M+N) / 2))
	exit := false
	for d := 1; d <= mid; d++ {
		fmt.Println("depth =", d)

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

		fmt.Println("F: ", F)

		// reverse search
		for k := -d + Rk; k <= d+Rk; k += 2 {
			var u int
			if k == d+Rk && (k != -d+Rk && R[k+1] < R[k-1]) {
				// move up
				u = R[k-1]
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
			fmt.Println("R: ", R)
			x, exists := F[k]
			if exists && x >= u {
				fmt.Println("found overlap")
				// fmt.Println(fmt.Sprintf("x=%x"))
				fmt.Println(fmt.Sprintf("x=%v, u=%v", x, u))
				// get the start and end of the snake
				y := x - k

				// solve left and right subproblems
				leftRect := boundingRect{left: rect.left, top: rect.top, right: u, bottom: v}
				rightRect := boundingRect{left: x, top: y, right: rect.right, bottom: rect.bottom}

				fmt.Println("leftRect: ", leftRect)
				fmt.Println("rightRect: ", rightRect)

				leftDiffs := myersDiffLinear(a, b, leftRect)
				fmt.Println("left diffs: ", leftDiffs)
				rightDiffs := myersDiffLinear(a, b, rightRect)
				fmt.Println("right diffs: ", rightDiffs)

				diffs = append(diffs, leftDiffs...)
				if x > u {
					// non empty snake
					diffs = append(diffs, Diff[T]{
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

		if exit {
			break
		}
	}

	return diffs
}
