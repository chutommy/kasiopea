package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type p struct {
	money int64
	extra int64
}

func main() {

	// create an output file
	f, err := os.Create("c.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over problems
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// store input
		px := make([]p, N)
		for n := 0; n < N; n++ {
			fmt.Scan(&px[n].money)
		}

		// solve
		s := solve(px)
		fmt.Fprintln(f, s)
	}
}

func solve(px []p) int64 {

	var total int64

	// sort px
	sort.SliceStable(px, func(i, j int) bool {
		return px[i].money > px[j].money
	})

	// simulate
	for !done(px) {

		// get data about last
		last := px[len(px)-1]
		lastMax := last.money + last.extra

		// divide by two until he is less or equal to the last
		count := px[0].extra
		currM := px[0].money
		for lastMax < currM {
			if currM%2 == 1 {
				count++
			}
			currM /= 2
		}

		// update data
		updated := p{currM, count}
		px = append(px[1:], updated)
		total += count
	}

	return total
}

func done(px []p) bool {
	last := px[len(px)-1]
	if l := last.money + last.extra; l >= px[0].money {
		return true
	}
	return false
}
