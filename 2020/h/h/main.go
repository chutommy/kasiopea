package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create an output file
	f, err := os.Create("h.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over problems
	for t := 0; t < T; t++ {

		// get R (rows), C (columns), N (budget, blocks)
		var N int
		R, C = 0, 0 // reset
		fmt.Scanf("%d %d %d", &R, &C, &N)

		// init new grids
		grid = make([][]int, R)
		for r := 0; r < R; r++ {
			grid[r] = make([]int, C)
		}
		durGrid = make([][]int, R)
		for r := 0; r < R; r++ {
			durGrid[r] = make([]int, C)
		}

		// store blocks backwards
		blocks := make([][2]int, N)
		for n := 0; n < N; n++ {

			// get r (row number), c (column number)
			var r, c int
			fmt.Scanf("%d %d", &r, &c)
			r--
			c--

			// store
			blocks[N-n-1] = [2]int{r, c}
		}

		// solve
		s := solve(blocks)
		fmt.Fprintln(f, N-s)
	}
}

var grid [][]int
var durGrid [][]int // durability grid

// R - rwos
var R int

// C - columns
var C int

func printGrid() {
	fmt.Println("=== GRID ===")
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("============")
}

func toGrid(blocks [][2]int) {

	// store blocks
	for _, coor := range blocks {
		grid[coor[0]][coor[1]]--
		durGrid[coor[0]][coor[1]]++
	}
}

func solve(blocks [][2]int) int {

	// fill the grid
	toGrid(blocks)

	// range over blocks and unlock them (one by one)
	for _, block := range blocks {

		// unlock
		if res := unlock(block[0], block[1]); res != -1 {
			return res
		}
	}

	//TODO remove this to increase performance
	// just for checking invalid situation
	// this should never panic
	panic("WTF 997 solve")
}

func unlock(r, c int) int {

	//TODO remove this to increase performance
	// just for checking invalid situation
	// this should never panic
	if grid[r][c] >= 0 {
		panic("WTF 102 unlock")
	}
	// ==========

	// increase tile value
	grid[r][c]++

	// check durability
	if grid[r][c] < 0 {
		return -1
	}

	// first row
	if r == 0 {
		grid[r][c] = 2
	} else {

		// get smallest positive adjs
		adjs := nextTo(r, c)
		min := -1
		for _, adj := range adjs {

			// get value
			adjVal := grid[adj[0]][adj[1]]

			// compare
			if adjVal > 0 {
				if min == -1 || adjVal < min {
					min = adjVal
				}
			}
		}

		// no valuable neighbour
		if min == -1 {
			return -1
		}

		// increase the value by the total durability - remove block
		grid[r][c] = min + durGrid[r][c]
	}

	// spread from here, if finish return result
	if res := spread(r, c); res != -1 {
		return res
	}
	return -1
}

// spread spreads the tile as much as possile
func spread(r, c int) int {

	// init que
	q := [][2]int{
		{r, c},
	}

	// get current value
	currVal := grid[r][c]

	//TODO remove this to increase performance
	// just for checking invalid situation
	// this should never panic
	if currVal <= 0 {
		panic("WTF 394 spread")
	}
	// ==========

	// until que is empty
	for len(q) > 0 {

		// deque
		tile := q[0]
		q = q[1:]

		// check for last row
		if tile[0] == R-1 {
			return currVal - 1
		}

		// get adjs
		adjs := nextTo(tile[0], tile[1])

		// range over adjs
		for _, adj := range adjs {

			adjVal := grid[adj[0]][adj[1]]

			// check value
			if adjVal != 0 && adjVal <= currVal {
				continue
			}

			// set
			grid[adj[0]][adj[1]] = currVal

			// add to que
			q = append(q, adj)
		}
	}

	return -1
}

// nextTo returns neighbour tiles
// include blocks
func nextTo(r, c int) [][2]int {

	// init neigbour tiles list
	var adj [][2]int

	// up
	if r != 0 {
		adj = append(adj, [2]int{r - 1, c})
	}

	// down
	if r != R-1 {
		adj = append(adj, [2]int{r + 1, c})
	}

	// left
	if c == 0 {
		// left border
		adj = append(adj, [2]int{r, C - 1})
	} else {
		adj = append(adj, [2]int{r, c - 1})
	}

	// right
	if c == C-1 {
		// right border
		adj = append(adj, [2]int{r, 0})
	} else {
		adj = append(adj, [2]int{r, c + 1})
	}

	return adj
}
