package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create output file
	f, err := os.Create("d.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		// get S, V, N
		var X, Y, N int
		fmt.Scan(&X, &Y, &N)

		// make grid
		p := &paper{
			X:    X,
			Y:    Y,
			grid: make([][]int, Y),
		}
		for y := 0; y < Y; y++ {
			p.grid[y] = make([]int, X)
		}

		// get init points
		q := []*coordinate{}
		for n := 0; n < N; n++ {

			// get coordinate
			var x, y int
			fmt.Scan(&x, &y)

			// store
			p.grid[y][x] = 1
			q = append(q, &coordinate{x, y})
		}

		// solve
		s := solve(q, p)
		fmt.Fprintln(f, s)
	}
}

type coordinate struct {
	x int
	y int
}

type paper struct {
	X    int
	Y    int
	grid [][]int
}

func solve(q []*coordinate, p *paper) int {

	time, count := 0, p.X*p.Y-len(q)

	for count > 0 {
		q, count = spreadFire(q, p, count)
		time++
	}

	return time
}

func spreadFire(q []*coordinate, p *paper, count int) ([]*coordinate, int) {

	// store all new fires into the newq
	newq := []*coordinate{}

	for _, coor := range q {

		xmin, xmax := coor.x == 0, coor.x == p.X-1
		ymin, ymax := coor.y == 0, coor.y == p.Y-1

		// spread to all directions if possible
		if !xmin {
			newq, count = setFire(coor.x-1, coor.y, newq, p, count)
		}
		if !xmax {
			newq, count = setFire(coor.x+1, coor.y, newq, p, count)
		}
		if !ymin {
			newq, count = setFire(coor.x, coor.y-1, newq, p, count)
		}
		if !ymax {
			newq, count = setFire(coor.x, coor.y+1, newq, p, count)
		}
		if !xmin && !ymin {
			newq, count = setFire(coor.x-1, coor.y-1, newq, p, count)
		}
		if !xmin && !ymax {
			newq, count = setFire(coor.x-1, coor.y+1, newq, p, count)
		}
		if !xmax && !ymin {
			newq, count = setFire(coor.x+1, coor.y-1, newq, p, count)
		}
		if !xmax && !ymax {
			newq, count = setFire(coor.x+1, coor.y+1, newq, p, count)
		}

		p.grid[coor.y][coor.x] = 2
	}

	return newq, count
}

func setFire(x, y int, newq []*coordinate, p *paper, count int) ([]*coordinate, int) {

	// if paper is ok, set it to fire
	if p.grid[y][x] == 0 {

		p.grid[y][x] = 1

		newq = append(newq, &coordinate{x, y})
		count--
	}

	return newq, count
}
