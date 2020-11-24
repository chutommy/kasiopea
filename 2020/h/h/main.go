package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
)

type coor [2]int

type block [3]int

type priorityQueue []block

func (p *priorityQueue) Len() int {
	return len(*p)
}

func (p *priorityQueue) Less(i int, j int) bool {
	return (*p)[i][2] < (*p)[j][2]
}

func (p *priorityQueue) Swap(i int, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *priorityQueue) Push(x interface{}) {
	*p = append(*p, x.(block))
}

func (p *priorityQueue) Pop() interface{} {
	lm := len(*p) - 1
	this := (*p)[lm]
	*p = (*p)[:lm]
	return this
}

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

		// get N (budget, blocks)
		var N int
		// reset R, C
		R, C = 0, 0
		fmt.Scanf("%d %d %d", &R, &C, &N)
		R++ // to add zero line

		// store blocks backwards
		blocks = make([]block, N)
		for n := 0; n < N; n++ {

			// store r (row number), c (column number)
			var r, c int
			fmt.Scanf("%d %d", &r, &c)
			blocks[N-n-1] = [3]int{r, c - 1, N - n}
		}

		// solve
		fmt.Fprintln(f, N-solve())

		// print status
		fmt.Printf("%d/%d done\n", t+1, T)
	}
}

var grid [][]int
var durGrid [][]int
var blockGrid [][][]block
var firstHitGrid [][]int
var visited [][]bool
var blocks []block

var pq *priorityQueue

// R rows
var R int

// C columns
var C int

func solve() int {

	// build structures
	buildGrids()

	// reset priorityQueue
	pq = &priorityQueue{}

	// init first row and set pq
	if r := bfs(1, coor{0, 0}); r != -1 {
		return 0
	}

	// query pq
	for len(*pq) > 0 {
		c := dequeue()
		if r := unlock(c); r != -1 {
			return r - 1
		}
	}

	panic("solution not found")
}

func buildGrids() {

	// make grids
	grid = make([][]int, R)
	durGrid = make([][]int, R)
	blockGrid = make([][][]block, R)
	firstHitGrid = make([][]int, R)
	visited = make([][]bool, R)

	// make subgrids
	for r := 0; r < R; r++ {
		grid[r] = make([]int, C)
		durGrid[r] = make([]int, C)
		blockGrid[r] = make([][]block, C)
		firstHitGrid[r] = make([]int, C)
		visited[r] = make([]bool, C)
	}

	// fill with blocks
	for _, c := range blocks {
		grid[c[0]][c[1]]--
		durGrid[c[0]][c[1]]++
		blockGrid[c[0]][c[1]] = append(blockGrid[c[0]][c[1]], c)
	}
}

func unlock(c coor) int {

	if len(blockGrid[c[0]][c[1]]) != 0 {
		// if tile is not destroyed enque the block under it
		enqueue(c)
		return -1
	}

	// bfs with the smallest adj + durability
	return bfs(firstHitGrid[c[0]][c[1]]+durGrid[c[0]][c[1]], c)
}

func enqueue(c coor) {

	// remove from blockGrid
	this := blockGrid[c[0]][c[1]][0]
	blockGrid[c[0]][c[1]] = blockGrid[c[0]][c[1]][1:]

	// enque into the pq
	heap.Push(pq, this)
}

func dequeue() coor {

	// get smallest block
	b := heap.Pop(pq).(block)

	return coor{b[0], b[1]}
}

func bfs(v int, sc coor) int {

	// init queue
	q := []coor{sc}

	// mark as visited
	visited[sc[0]][sc[1]] = true

	for len(q) > 0 {

		// dequeue
		c := q[0]
		q = q[1:]

		// v == 0 (true)
		// mark with value
		grid[c[0]][c[1]] = v

		// check for last row
		if c[0] == R-1 {
			return v
		}

		// enqueue adj
		for _, adj := range getAdjs(c[0], c[1]) {

			// check value
			val := grid[adj[0]][adj[1]]
			if val < 0 {

				// if has more durability
				if len(blockGrid[adj[0]][adj[1]]) != 0 {
					firstHitGrid[adj[0]][adj[1]] = v
					enqueue(adj)
				}
				continue

			} else if val != 0 || visited[adj[0]][adj[1]] {
				continue
			}

			// append and mark as visited
			q = append(q, adj)
			visited[adj[0]][adj[1]] = true
		}
	}

	// bfs did not end in the last row
	return -1
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
