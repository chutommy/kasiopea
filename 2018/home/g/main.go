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

	// get T (problems)
	var T int
	fmt.Scan(&T)

	// range over problems
	for t := 0; t < T; t++ {

		// get N (vertices), M (edges)
		var N, M int
		fmt.Scanf("%d %d", &N, &M)

		// create an adjacency list and edge list
		adj := make([][]int, N)
		edges := make([][3]int, M)

		// process edges
		for m := 0; m < M; m++ {

			// get an edge
			var U, V, P int
			fmt.Scan("%d %d %d", &U, &V, &P)

			// store data
			adj[U] = append(adj[U], V)
			edges = append(edges, [3]int{U, V, P})
		}

		// solve
		r := solve(adj, edges)
		fmt.Fprintln(f, r)
	}
}

func solve(adj [][]int, edges [][3]int) int {

	// TODO

	return -1
}
