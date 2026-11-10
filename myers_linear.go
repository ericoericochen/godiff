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

// diagonal along the edit graph where the elements match
// u - v = x - y
type myersSnake struct {
	// start coordinate of snake
	u int
	v int

	// end coordinate of snake
	x int
	y int
}

func (s myersSnake) empty() bool {
	return s.x == s.u && s.y == s.v
}

func MyersDiffLinear(a, b []string) []Diff {
	rect := boundingRect{
		left:   0,
		top:    0,
		right:  len(a),
		bottom: len(b),
	}
	return myersDiffLinear(a, b, rect)
}

func myersDiffLinear(a, b []string, rect boundingRect) []Diff {
	fmt.Println("Myers Diff Linear: ", rect)

	// base case
	M := rect.right - rect.left
	N := rect.bottom - rect.top
	diffs := []Diff{}

	// empty rect = no diffs
	if M == 0 && N == 0 {
		return diffs
	}

	// deletion
	if M > 0 && N == 0 {
		diffs = append(diffs, Diff{Type: Delete, A1: rect.left, A2: rect.right})
		return diffs
	}

	// insertion
	if N > 0 && M == 0 {
		diffs = append(diffs, Diff{Type: Insert, B1: rect.top, B2: rect.bottom})
		return diffs
	}

	// find the middle snake in the edit graph
	snake := findMiddleSnake(a, b, rect)

	// solve left and right subproblems
	leftRect := boundingRect{left: rect.left, top: rect.top, right: snake.u, bottom: snake.v}
	rightRect := boundingRect{left: snake.x, top: snake.y, right: rect.right, bottom: rect.bottom}

	leftDiffs := myersDiffLinear(a, b, leftRect)
	rightDiffs := myersDiffLinear(a, b, rightRect)
	diffs = append(diffs, leftDiffs...)
	// non empty snakes are equal sequences
	if !snake.empty() {
		diff := Diff{
			Type: Equal,
			A1:   snake.u,
			A2:   snake.x,
		}
		diffs = append(diffs, diff)
	}
	diffs = append(diffs, rightDiffs...)
	return diffs
}

// Find the middle snake in the edit graph bounded by rect.
// Performs a forward and backward search to find the overlapping snake.
// The endpoints of the snake (x, y) and (u, v) lie on the same diagonal
// such that x - y = u - v and x >= u.
func findMiddleSnake(a, b []string, rect boundingRect) myersSnake {
	// store the frontier of forward and reverse search
	fk := rect.left - rect.top
	rk := rect.right - rect.bottom
	f := map[int]int{fk + 1: rect.left}
	r := map[int]int{rk - 1: rect.right}

	M := rect.right - rect.left
	N := rect.bottom - rect.top
	mid := int(math.Ceil(float64(M+N) / 2))

	for d := 0; d <= mid; d++ {
		forwardSearch(a, b, d, fk, f, r, rect)
		snake := reverseSearch(a, b, d, rk, f, r, rect)
		if snake != nil {
			return *snake
		}
	}

	return myersSnake{}
}

// Performs forward search and returns the overlapping snake if found
// a - source string sequence
// b - dest string sequence
// d - current depth of search
// fk - starting k at depth = 0
// f - frontier of forward search storing (k, x)
// r - frontier of reverse search storing (k, x)
// rect - bounding rect of search space
func forwardSearch(a []string, b []string, d int, fk int, f map[int]int, r map[int]int, rect boundingRect) *myersSnake {
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
		if exists && x >= u {
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

// Perform the reverse search and returns the overlapping snake if found
// a - source string sequence
// b - dest string sequence
// d - current depth of search
// rk - starting k at depth = 0
// f - frontier of forward search storing (k, x)
// r - frontier of reverse search storing (k, x)
// rect - bounding rect of search space
func reverseSearch(a []string, b []string, d int, rk int, f map[int]int, r map[int]int, rect boundingRect) *myersSnake {
	for k := d + rk; k >= -d+rk; k -= 2 {
		var u int
		if k == d+rk || (k != -d+rk && r[k+1] < r[k-1]) {
			// move up
			u = r[k-1]
		} else {
			// move left
			u = r[k+1] - 1
		}
		v := u - k
		for u > rect.left && v > rect.top && a[u-1] == b[v-1] {
			u--
			v--
		}
		r[k] = u
		x, exists := f[k]
		if exists && x >= u {
			y := x - k
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
