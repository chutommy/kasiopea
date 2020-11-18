package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create an output file
	f, err := os.Create("a.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over each problem
	for t := 0; t < T; t++ {

		// get N, P
		var N, P int
		fmt.Scan(&N, &P)

		// create map to store data
		// values are m[hated]= []int{beloved, speed}
		m := make(map[int][]int)

		// range over seniors
		for n := 0; n < N; n++ {

			// get data of the senior
			var a, b, c int
			fmt.Scanf("%d %d %d", &a, &b, &c)

			// store
			if d, ok := m[a]; ok {

				// check speed
				if c > d[1] {
					m[a] = []int{b, c}
				}

			} else {
				// first occurence
				m[a] = []int{b, c}
			}
		}

		// solve
		s := solve(P, m)
		fmt.Fprintln(f, s)
	}
}

func solve(P int, m map[int][]int) int {

	// current
	curr := P

	// visited
	visited := make(map[int]struct{})
	visited[P] = struct{}{}

	// simulate
	for {

		// next
		if v, ok := m[curr]; ok {
			curr = v[0]
		} else {
			return curr
		}

		// check for infinite loop
		if _, ok := visited[curr]; ok {
			return -1
		}

		// mark current as visited
		visited[curr] = struct{}{}
	}
}
