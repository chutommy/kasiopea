package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create an output file
	f, err := os.Create("f.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over each problem
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// init adjacency list
		adj := make([]int, N)

		// store integers
		for n := 0; n < N; n++ {

			// get int
			var i int
			fmt.Scan(&i)

			// store vector
			v := n + i
			if v >= N {
				v -= N
				if v > n {
					adj[n] = -1
					continue
				}
			}
			adj[n] = v
		}

		// solve
		s := solve(N, adj)
		if s {
			fmt.Fprintln(f, "ANO")
		} else {
			fmt.Fprintln(f, "NE")
		}
		fmt.Println(t)
	}
}

func solve(N int, adj []int) bool {

	// init status arrays
	visited := make([]bool, N)
	// visitedCopy := make([]bool, N)
	rooted := make([]int, N)
	// rootedCopy := make([]int, N)

	// dfs
	for root, next := range adj {

		// if visited skip
		if visited[root] {
			continue
		}

		visited[root] = true
		// copy(visitedCopy, visited)

		rooted[root] = root
		// copy(rootedCopy, rooted)

		cycle := false
		dropped := false

		// traversal
		for {

			// check if valid
			if next == -1 {
				break
			}

			// check if visited
			if visited[next] {

				// check cycle
				if rooted[next] == root {
					cycle = true
				}
				break
			}
			visited[next] = true
			rooted[next] = root

			// to next
			nextAdj := adj[next]
			if nextAdj < next {
				if dropped {
					// copy(visited, visitedCopy)
					// copy(rooted, rootedCopy)
					break
				} else {
					dropped = true
				}
			}
			next = nextAdj
		}

		// return true if cycle found
		if cycle {
			return true
		}
	}

	return false
}
