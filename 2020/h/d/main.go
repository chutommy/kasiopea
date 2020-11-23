package main

import (
	"fmt"
	"log"
	"os"
)

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

	// range over problems
	for t := 0; t < T; t++ {

		// get N, M
		var N, M int
		fmt.Scanf("%d %d", &N, &M)

		// init DT
		stats = make([]int, N)
		visited = make([]bool, N)
		friends := make([][]int, N)
		for n := 0; n < N; n++ {
			friends[n] = []int{}
		}

		// store graph
		for m := 0; m < M; m++ {

			// get u, v
			var u, v int
			fmt.Scanf("%d %d", &u, &v)
			u--
			v--

			// connect
			friends[u] = append(friends[u], v)
			friends[v] = append(friends[v], u)
		}

		// solve
		s := solve(N, M, friends)
		if s {
			fmt.Fprintln(f, "ANO")

			// range over stats
			var ones []int
			for i, v := range stats {
				if v == 1 {
					ones = append(ones, i+1)
				}
			}

			// print solution
			fmt.Fprintln(f, len(ones))
			for _, v := range ones {
				fmt.Fprintln(f, v)
			}

		} else {
			fmt.Fprintln(f, "NE")
		}

		// print status
		fmt.Printf("%d/%d done\n", t+1, T)
	}
}

// 0 - not yet
// 1 - pravitko
// 2 - kruzitko
var stats []int
var visited []bool

func solve(N, M int, friends [][]int) bool {

	// check - everybody has at least one friend
	for n := 0; n < N; n++ {
		if len(friends[n]) == 0 {
			return false
		}
	}

	// set each pupil
	for i := 0; i < N; i++ {

		if stats[i] != 0 {
			continue
		}

		bfs(N, i, friends)
	}

	return true
}

func bfs(N, i int, friends [][]int) {

	// init que
	q := []int{i}
	visited[i] = true

	stats[i] = 1

	// until que is empty
	for len(q) > 0 {

		// deque
		this := q[0]
		q = q[1:]

		// get curr status
		toSet := change(stats[this])

		// range over neighbours
		for _, next := range friends[this] {
			if !visited[next] {
				q = append(q, next)
				stats[next] = toSet
				visited[next] = true
			}
		}
	}
}

func change(i int) int {
	if i == 1 {
		return 2
	}
	return 1
}
