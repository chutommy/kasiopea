package main

import (
	"fmt"
	"log"
	"os"
)

type coor struct {
	x int
	y int
}

func main() {

	// create an output file
	f, err := os.Create("d.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over each problem
	for t := 0; t < T; t++ {

		// get S, V, N
		var S, V, N int
		fmt.Scanf("%d %d %d", &S, &V, &N)

		// get sources
		srcs := make([]coor, N)
		for n := 0; n < N; n++ {

			// get coordinate
			var x, y int
			fmt.Scanf("%d %d", &x, &y)

			// store coordinate
			srcs[n] = coor{
				x: x,
				y: y,
			}
		}

		// solve
		s := solve(S, V, N, srcs)
		fmt.Fprintln(f, s)
		fmt.Printf("%d/%d done\n", t+1, T)
	}
}

func solve(S, V, N int, srcs []coor) int {

	// create a grid
	grid := make([][]int, V)
	for v := 0; v < V; v++ {
		grid[v] = make([]int, S)
	}

	// set fire sources
	for _, c := range srcs {
		grid[c.y][c.x] = 1
	}

	// burn paper
	count := -1
	for len(srcs) > 0 {
		count++

		// create a new list of sources
		newsrcs := &[]coor{}

		// spread fire
		for _, c := range srcs {

			minx, maxx := c.x == 0, c.x == S-1
			miny, maxy := c.y == 0, c.y == V-1

			// set sides
			if !minx {
				setFire(coor{c.x - 1, c.y}, newsrcs, grid)
			}
			if !maxx {
				setFire(coor{c.x + 1, c.y}, newsrcs, grid)
			}
			if !miny {
				setFire(coor{c.x, c.y - 1}, newsrcs, grid)
			}
			if !maxy {
				setFire(coor{c.x, c.y + 1}, newsrcs, grid)
			}

			// set diagonals
			if !minx && !miny {
				setFire(coor{c.x - 1, c.y - 1}, newsrcs, grid)
			}
			if !minx && !maxy {
				setFire(coor{c.x - 1, c.y + 1}, newsrcs, grid)
			}
			if !maxx && !miny {
				setFire(coor{c.x + 1, c.y - 1}, newsrcs, grid)
			}
			if !maxx && !maxy {
				setFire(coor{c.x + 1, c.y + 1}, newsrcs, grid)
			}
		}

		srcs = *newsrcs
	}

	return count
}

func setFire(c coor, srcs *[]coor, grid [][]int) {

	// check
	if grid[c.y][c.x] == 0 {

		// set fire
		grid[c.y][c.x] = 1

		// add to the source list
		*srcs = append(*srcs, coor{
			x: c.x,
			y: c.y,
		})
	}
}
