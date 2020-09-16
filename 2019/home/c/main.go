package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create an output file
	f, err := os.Create("c.out")
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

		// get path
		path := make([]int, N)
		for n := 0; n < N; n++ {

			// get i
			var i int
			fmt.Scan(&i)

			path[n] = i
		}

		res := solve(path, N)
		fmt.Fprintln(f, res)
	}
}

func solve(path []int, N int) int {

	// define current pos and counter
	var pos, count int

	// until finish
	for pos < N {

		var best, inx int
		// find the best jump opt
		for i := 1; i <= path[pos]; i++ {

			// over/in finish
			v := pos + i
			if v >= N {
				pos = N
				break
			}

			// compare the best
			sum := i + path[v]
			if sum > best {
				best = sum
				inx = i
			}
		}

		pos += inx
		count++
	}

	return count
}
