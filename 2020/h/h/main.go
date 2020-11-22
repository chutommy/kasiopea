package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

		// store blocks backwards
		blocks = make([][2]int, N)
		for n := 0; n < N; n++ {

			// get r (row number), c (column number)
			var r, c int
			fmt.Scanf("%d %d", &r, &c)
			r--
			c--

			// store
			blocks[N-n-1] = [2]int{r, c}
		}

		// init data types
		buildGrids()
		buildPrior()
		initPQ()
		initFromTop()

		// solve
		s := solve(N)
		fmt.Fprintln(f, s)

		fmt.Printf("%d/%d done\n", t+1, T)
	}
}

// -x durability
// 0 none
// 1 processing
// 2 processed
var grid [][]int
var durGrid [][]int
var dessGrid [][]int
var backGrid [][]int
var pq [][2]int
var prior [][][]int
var blocks [][2]int
var R, C int

func solve(N int) int {

	// do magic
	pos := transform()
	backtrack(pos)

	// printGrid()
	// printDurGrid()
	// printDessGrid()
	// printPrior()
	// printBackGrid()

	// find min
	min := -1
	for _, c := range backGrid[0] {
		if c != 0 && (min == -1 || c < min) {
			min = c
		}
	}

	return N - min
}

func transform() int {

	// until solution is found
	for len(pq) > 0 {
		next := deque()
		if r := unlock(next[0], next[1]); r != -1 {
			return r
		}
	}

	panic("WTF solution not found")
}

func printGrid() {
	fmt.Println("=== GRID ===")
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("============")
}

func printDurGrid() {
	fmt.Println("=== DURGRID ===")
	for _, row := range durGrid {
		fmt.Println(row)
	}
	fmt.Println("===============")
}

func printDessGrid() {
	fmt.Println("=== DESSGRID ===")
	for _, row := range dessGrid {
		fmt.Println(row)
	}
	fmt.Println("===============")
}

func printBackGrid() {
	fmt.Println("=== BACKGRID ===")
	for _, row := range backGrid {
		fmt.Println(row)
	}
	fmt.Println("===============")
}

func buildGrids() {

	// create grid
	grid = make([][]int, R)
	for r := 0; r < R; r++ {
		grid[r] = make([]int, C)
	}

	// create durGrid
	durGrid = make([][]int, R)
	for r := 0; r < R; r++ {
		durGrid[r] = make([]int, C)
	}

	// create dessGrid
	dessGrid = make([][]int, R)
	for r := 0; r < R; r++ {
		dessGrid[r] = make([]int, C)
	}

	// create backGrid
	backGrid = make([][]int, R)
	for r := 0; r < R; r++ {
		backGrid[r] = make([]int, C)
	}

	// fill with the blocks
	for _, coor := range blocks {
		grid[coor[0]][coor[1]]--
		durGrid[coor[0]][coor[1]]++
		dessGrid[coor[0]][coor[1]]++
	}
}

func printPrior() {
	fmt.Println("=== PRIOR ===")
	for _, row := range prior {
		fmt.Println(row)
	}
	fmt.Println("=============")
}

func buildPrior() {

	// create matrice
	prior = make([][][]int, R)
	for r := 0; r < R; r++ {
		prior[r] = make([][]int, C)
		for c := 0; c < C; c++ {
			prior[r][c] = []int{}
		}
	}

	// add priorities
	for p, coor := range blocks {
		prior[coor[0]][coor[1]] = append(prior[coor[0]][coor[1]], p)
	}
}

func initPQ() {

	// create priority que
	pq = [][2]int{}

	// add first row
	for c := 0; c < C; c++ {
		if grid[0][c] == -1 {
			enque([2]int{0, c})
		}
	}
}

func enque(coor [2]int) {

	// decrease durability
	if durGrid[coor[0]][coor[1]] == 0 {
		return
	}
	durGrid[coor[0]][coor[1]]--

	val := prior[coor[0]][coor[1]][0]

	// find index to insert
	pos := sort.Search(val, func(i int) bool {
		if len(pq) <= i {
			return true
		}
		this := pq[i]
		return val < prior[this[0]][this[1]][0]
	})

	// remove from prior add to pq
	pq = append(pq[:pos], append([][2]int{coor}, pq[pos:]...)...)
}

func deque() [2]int {

	// select first and remove from prior
	next := pq[0]
	prior[next[0]][next[1]] = prior[next[0]][next[1]][1:]

	// enque next
	if len(prior[next[0]][next[1]]) != 0 {
		enque(next)
	}

	pq = pq[1:]
	return next
}

func initFromTop() int {

	// range over first row
	for c := 0; c < C; c++ {
		if grid[0][c] == 0 {
			if pos := spread(0, c); pos != -1 {
				return pos
			}
		}
	}

	return -1
}

func unlock(r, c int) int {

	// decrease durability
	grid[r][c]++

	// if block destroyed
	if grid[r][c] == 0 {

		// check adjs
		if r == 0 {
			return spread(0, c)
		} else {
			for _, adj := range getAdjs(r, c) {
				if grid[adj[0]][adj[1]] == 2 {
					return spread(adj[0], adj[1])
				}
			}
		}
	}

	return -1
}

func spread(r, c int) int {

	// init que
	q := [][2]int{
		{r, c},
	}

	// until que is empty
	for len(q) > 0 {

		// deque
		tile := q[0]
		q = q[1:]

		// mark
		grid[tile[0]][tile[1]] = 2

		// check for last row
		if tile[0] == R-1 {
			return tile[1]
		}

		// get neighbours
		adjs := getAdjs(tile[0], tile[1])

		// range over adjs
		for _, adj := range adjs {

			adjVal := grid[adj[0]][adj[1]]

			// check status
			if adjVal != 0 {
				if adjVal < 0 {
					enque(adj)
				}
				continue
			}

			// set
			grid[adj[0]][adj[1]] = 1

			// add to the que
			q = append(q, adj)
		}
	}

	return -1
}

func backtrack(pos int) {

	// init que
	q := [][2]int{
		{R - 1, pos},
	}

	// base value for grid's starting point
	backGrid[R-1][pos] = dessGrid[0][pos]

	// until que is empty
	for len(q) > 0 {

		// deque
		tile := q[0]
		q = q[1:]

		currVal := backGrid[tile[0]][tile[1]]

		// get neighbours
		adjs := getAdjs(tile[0], tile[1])

		// range over adjs
		for _, adj := range adjs {

			adjVal := grid[adj[0]][adj[1]]

			// check status
			if adjVal != 2 {
				continue
			}

			// set
			bv := backGrid[adj[0]][adj[1]]
			sv := dessGrid[adj[0]][adj[1]] + currVal
			if bv == 0 || bv > sv {
				backGrid[adj[0]][adj[1]] = sv
				q = append(q, adj)
			}
		}
	}
}

// getAdjs returns neighbour tiles
// include blocks
func getAdjs(r, c int) [][2]int {

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
