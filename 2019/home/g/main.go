package main

import (
	"fmt"
	"log"
	"os"
)

type poet struct {
	name  int
	liked []*poet
}

func main() {

	// create an output file
	f, err := os.Create("g.out")
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

		// init poets
		poets := make([]*poet, N)
		for n := 0; n < N; n++ {
			poets[n] = &poet{}
		}

		// get poets
		for m := 0; m < M; m++ {

			var a, b int
			fmt.Scanf("%d %d", &a, &b)
			a--
			b--

			poets[b].liked = append(poets[b].liked, poets[a])
		}

		s := solve(poets, N)
		fmt.Fprintln(f, s)
	}
}

func solve(poets []*poet, N int) int {

	var root *poet
	// find root
	for n := 0; n < N; n++ {
		if p := poets[n]; len(p.liked) == 0 {
			root = p
			break
		}
	}

	q := []*poet{root}
	var layer, inlayer, innextlayer = 1, 1, 0
	for len(q) != 0 {

		p := q[0]
		q = q[1:]

		for _, add := range p.liked {
			q = append(q, add)
			innextlayer++
			fmt.Println("inlay", inlayer)
		}

		inlayer--
		if inlayer == 0 {
			layer++
			inlayer = innextlayer
			innextlayer = 0
		}
	}

	return layer
}
