package main

import (
	"fmt"
	"log"
	"os"
)

type bulb struct {
	status int
	// 0 - unvisited
	// 1 - in queue (processing)
	// 2 - visited
	color int
	wired []*bulb
}

func main() {

	// create output file
	f, err := os.Create("d.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		// get N, M
		var N, M int
		fmt.Scan(&N, &M)

		// init bulbs
		bulbs := make([]*bulb, N)
		for n := 0; n < N; n++ {
			bulbs[n] = &bulb{}
		}

		// wire bulbs
		for m := 0; m < M; m++ {

			// get u, v
			var u, v int
			fmt.Scan(&u, &v)

			// store
			bulbs[u-1].wired = append(bulbs[u-1].wired, bulbs[v-1])
			bulbs[v-1].wired = append(bulbs[v-1].wired, bulbs[u-1])
		}

		// assign colors
		for n := 0; n < N; n++ {

			// get color
			var c int
			fmt.Scan(&c)

			// assign color
			bulbs[n].color = c
		}

		// solve
		s := solve(N, bulbs)
		fmt.Fprintln(f, s)
	}
}

func solve(N int, bulbs []*bulb) int {
	var total int

	for n := 0; n < N; n++ {
		b := bulbs[n]

		// check if not visited
		if b.status == 2 {
			continue
		}

		// count colors
		colors := bfs(b)

		// process colors
		var sum, top int
		for _, v := range colors {
			sum += v
			if v > top {
				top = v
			}
		}
		total += sum - top
	}

	return total
}

func bfs(b *bulb) map[int]int {

	colors := make(map[int]int)

	// add first
	q := []*bulb{b}

	for len(q) != 0 {

		// dequeue
		curr := q[0]
		curr.status = 2
		colors[curr.color]++
		q = q[1:]

		// range over wired
		for _, next := range curr.wired {

			// continue if processed or in queue
			if next.status != 0 {
				continue
			}

			// add to queue
			q = append(q, next)
			next.status = 1
		}
	}

	return colors
}
