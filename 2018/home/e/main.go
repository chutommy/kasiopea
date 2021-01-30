package main

import (
	"fmt"
	"log"
	"os"
)

type town struct {
	population    int
	distance      int     // distance to finals
	sumPopulation int     // how many people go through this town
	income        []*town // towns that point to this town
	edges         []*town
}

func main() {

	// create output file
	f, err := os.Create("e.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		// get N, M
		// N - towns
		// M - edges
		var N, M int
		fmt.Scan(&N, &M)

		// init towns
		towns := make([]*town, N)

		// iterate over edges
		for m := 0; m < M; m++ {

			// get u,v
			var u, v int
			fmt.Scan(&u, &v)

			// create edge
			towns[u].edges = append(towns[u].edges, towns[v])
			towns[v].edges = append(towns[v].edges, towns[u])
		}

		// iterate over pointers
		for n := 1; n < N; n++ {

			// get population and poionter
			var pop, point int
			fmt.Scan(&pop, &point)

			// store
			t := towns[n]
			t.population = pop
			towns[point].income = append(towns[point].income, t)
		}

		// solve
		u, v := solve()
		sol := fmt.Sprintf("%d %d", u, v)
		fmt.Fprintln(f, sol)
	}
}

func solve() (int, int) {

	// TODO

	return 0, 0
}
