package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type peak struct {
	x int
	h int
}

func main() {

	// create output file
	f, err := os.Create("e.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over problems
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		peaks := make([]*peak, N)
		// store peaks
		for n := 0; n < N; n++ {

			// get coordinate
			var x, h int
			fmt.Scanf("%d %d", &x, &h)

			// store
			peaks[n] = &peak{x, h}
		}

		// solve
		solution := solve(peaks, N)

		// format the result
		var result []string
		for i := 0; i < len(solution); i++ {
			result = append(result, fmt.Sprint(solution[i]))
		}
		s := strings.Join(result, " ")

		fmt.Fprintln(f, s)
	}
}

func solve(peaks []*peak, N int) []int {

	result := make([]int, N)

	// range over peaks
	for idx, peak := range peaks {

		var left, right int
		var leftB, rightB bool

		// search left
		if idx > 0 {
			for l := idx - 1; l >= 0; l-- {
				if peaks[l].h >= peak.h {

					left = dist(peak, peaks[l])
					leftB = true
					break
				}
			}
		}

		// search right
		if idx < N-1 {
			for l := idx + 1; l < N; l++ {
				if peaks[l].h >= peak.h {

					right = dist(peak, peaks[l])
					rightB = true
					break
				}
			}
		}

		// store
		switch {
		case leftB && rightB:
			if left < right {
				result[idx] = left
			} else {
				result[idx] = right
			}

		case leftB:
			result[idx] = left

		case rightB:
			result[idx] = right

		default:
			result[idx] = -1
		}
	}

	return result
}

func dist(p1, p2 *peak) int {

	// get diffs
	hDiff := p2.h - p1.h
	if hDiff < 0 {
		hDiff *= -1
	}

	xDiff := p2.x - p1.x
	if xDiff < 0 {
		xDiff *= -1
	}

	// calculate the distance
	return xDiff - hDiff
}
