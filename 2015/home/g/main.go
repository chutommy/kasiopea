package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create an output file
	f, err := os.Create("g.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over each problem
	for t := 0; t < T; t++ {

		// get N, M
		var N, M int
		fmt.Scanf("%d %d", &N, &M)

		// store grid
		grid := make([][]int, N)
		for n := 0; n < N; n++ {

			grid[n] = make([]int, M)

			// store heights
			for m := 0; m < M; m++ {

				// get integer
				var i int
				fmt.Scan(&i)

				// store integer
				grid[n][m] = i
			}
		}

		// solve
		s := solve(N, M, grid)
		fmt.Println(s)
	}
}

func solve(N, M int, grid [][]int) int {
	// get xgrid
	xgrid := make([][]int, N)
	for n := 0; n < N; n++ {

		// calculate value of the row
		xgrid[n] = solveRow(M, grid[n])
	}

	fmt.Println(xgrid)

	// get ygrid
	// merge xgrid and ygrid

	return 0
}

func solveRow(l int, row []int) []int {

	// minimal row
	if l < 3 {
		return make([]int, l)
	}

	// initialize rows
	arow := make([]int, l)
	brow := make([]int, l)

	// calculate new row
	a, b := 0, l-1
	amax, bmax := 0, 0
	// range over row
	for a != l {

		// set index a
		if v := row[a]; v > amax {
			arow[a] = 0
			amax = v
		} else {
			arow[a] = amax - v
		}

		// set index b
		if v := row[b]; v > bmax {
			brow[b] = 0
			bmax = v
		} else {
			brow[b] = bmax - v
		}

		a++
		b--
	}

	// merge rows
	newrow := make([]int, l)
	for i := 0; i < l; i++ {

		a := arow[i]
		b := brow[i]

		// select less
		if a < b {
			newrow[i] = a
		} else {
			newrow[i] = b
		}
	}

	return newrow
}
