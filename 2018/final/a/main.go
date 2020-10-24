package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// create output file
	f, err := os.Create("a.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over T problems
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// retrieve data
		data := make([]int, N)
		for n := 0; n < N; n++ {
			// get integer
			var i int
			fmt.Scan(&i)
			// store
			data[n] = i
		}

		// solve
		sol := solve(data)

		// print solution
		fmt.Fprintln(f, sol)
	}
}

func solve(d []int) string {
	l := len(d)

	// skip if impossible
	if l < 3 {
		return "-1"
	}

	// declare operating values
	a, b := d[0], d[1]
	unic := a == b
	cv, cp := -1, -1

	// range over data
	for i := 2; i < l; i++ {
		if unic {
			if d[i] == a {
				// a,b,c all would be same
				continue
			}
		}
		// found ok
		cv = d[i]
		cp = i
		break
	}

	if cv == -1 {
		return "-1"
	}

	// reorder
	if a == b {
		d[1], d[cp] = d[cp], d[1]
	} else {
		// find highest
		switch {
		case a > b && a > cv:
			d[0], d[1] = d[1], d[0]
		case cv > a && cv > b:
			d[1], d[cp] = d[cp], d[1]
		}
	}

	// join result into string
	resStr := make([]string, l)
	for p, v := range d {
		resStr[p] = fmt.Sprint(v)
	}
	res := strings.Join(resStr, " ")

	return res
}
