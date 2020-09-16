package main

import (
	"fmt"
	"log"
	"os"
)

type town struct {
	population int
	neighbours []*town
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

	for t := 0; t < T; t++ {

		// get N, M, K
		var N, M, K int
		fmt.Scanf("%d %d %d", &N, &M, &K)

		// get populations
		towns := make([]*town, N)
		for n := 0; n < N; n++ {

			t := &town{}
			fmt.Scan(&t.population)
			towns[n] = t
		}

		// get neighbours
		for m := 0; m < M; m++ {

			// get towns' indexes
			var u, v int
			fmt.Scanf("%d %d", &u, &v)
			u--
			v--

			// append
			towns[u].neighbours = append(towns[u].neighbours, towns[v])
			towns[v].neighbours = append(towns[v].neighbours, towns[u])
		}

		s := solve(towns, N, K)
		fmt.Fprintln(f, s)
	}
}

func solve(towns []*town, N, K int) string {

	// range over towns
	for n := 0; n < N; n++ {

		t := towns[n]

		npop := make(map[int]bool)

		// range over neighbours
		for _, n := range t.neighbours {
			if n.population > t.population {
				npop[n.population] = true
			}
		}

		// check
		k := K - t.population
		for i := 1; i <= k; i++ {
			if exist := npop[t.population+i]; !exist {
				return "ANO"
			}
		}
	}

	return "NE"
}
