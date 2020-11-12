package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {

	// create an output file
	f, err := os.Create("b.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over each problem
	for t := 0; t < T; t++ {

		// get N
		var N int
		fmt.Scan(&N)

		// get result
		res := solve(N)
		fmt.Fprintln(f, res)
	}
}

func solve(N int) string {
	s := decomp(N)
	return s
}

func decomp(N int) string {

	switch N {
	case 0:
		return ""
	case 1:
		return "1"
	}

	// divide by 3
	tx, n := 0, N
	for {
		next := n / 3

		// check
		if next == 0 {
			break
		} else {
			n = next
			tx++
		}
	}

	// init output string
	var out string
	// add powers of three
	for t := 0; t < tx; t++ {
		out += "(1+1+1)*"
	}

	total := int(math.Pow(float64(3), float64(tx)))
	// divide by two if possible
	if res := total * 2; res <= N {
		out += "(1+1)*"
		total = res
	}

	// remove last character
	out = out[:len(out)-1]

	// return solution
	if d := decomp(N - total); d != "" {
		out += "+" + d
	}

	return out
}
