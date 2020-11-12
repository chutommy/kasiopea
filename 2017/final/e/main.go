package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type stick struct {
	l int
	c int
}

func main() {

	// create an outpuf file
	f, err := os.Create("e.out")
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

		// retrieve sticks
		sticks := make([]stick, N)
		for n := 0; n < N; n++ {

			// get stats
			var l, c int
			fmt.Scan(&l, &c)

			// store
			sticks[n] = stick{l, c}
		}

		// get solution
		sol := solve(sticks)
		if sol {
			fmt.Fprintln(f, "ANO")
		} else {
			fmt.Fprintln(f, "NE")
		}
		fmt.Printf("%d/%d DONE\n", t+1, T)
	}
}

func solve(sticks []stick) bool {
	l := len(sticks)

	// sort sticks
	sort.Slice(sticks, func(i, j int) bool {
		return sticks[i].l < sticks[j].l
	})

	// initialize a,b,c
	a, b, c := l-1, l-1, l-1

	ca := sticks[a].c
	cb := sticks[b].c
	cc := sticks[c].c

	// move b
	for b > 0 && cb == ca {
		b--
		cb = sticks[b].c
	}
	// move c
	for c > 0 && (cc == ca || cc == cb) {
		c--
		cc = sticks[c].c
	}

	// query sticks
	for {
		// define current colors
		ca = sticks[a].c
		cb = sticks[b].c
		cc = sticks[c].c

		// move b
		for b > c && cb == ca {
			b--
			cb = sticks[b].c
		}
		// move c
		for c > 0 && (cc == ca || cc == cb) {
			c--
			cc = sticks[c].c
		}
		if cc == ca || cc == cb {
			return false
		}

		if sticks[b].l+sticks[c].l > sticks[a].l {
			return true
		}
		a--
	}
}
