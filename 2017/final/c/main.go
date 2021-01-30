package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create an output file
	f, err := os.Create("c.out")
	if err != nil {
		log.Fatal(err)
	}

	// get T
	var T int
	fmt.Scan(&T)

	// iterate over problems
	for t := 0; t < T; t++ {

		// get K
		var K int
		fmt.Scan(&K)

		// get N row
		var N int
		fmt.Scan(&N)

		// store first row
		fRow := make([]byte, N)
		for n := 0; n < N; n++ {
			fmt.Scanf("%c", &fRow[n])
		}
		fmt.Scanf("\n")

		// store rows
		rows := make([][]byte, K)
		for k := 0; k < K; k++ {

			// get Ni
			var Ni int
			fmt.Scan(&Ni)

			// store
			row := make([]byte, Ni)
			for ni := 0; ni < Ni; ni++ {
				fmt.Scanf("%c", &row[ni])
			}
			fmt.Scanf("\n")

			rows[k] = row
		}

		// solve
		s := solve(fRow, rows)
		fmt.Fprintln(f, s)
	}
}

func solve(fRow []byte, rows [][]byte) int {
	ll := len(fRow)
	lr := len(rows)

	// init final lengths array
	final := make([]int, lr)

	// init first queue
	var fq []byte
	for i := 0; i < ll; i++ {
		lastInd := len(fq) - 1

		// check for same values
		if lastInd != -1 && fq[lastInd] == fRow[i] {
			fq = fq[:lastInd]
		} else {
			fq = append(fq, fRow[i])
		}
	}

	// iterate over rows
	for ind, row := range rows {

		// init queue
		q := make([]byte, len(fq))
		copy(q, fq)

		// iterate over the row and simulate
		for _, b := range row {
			lastInd := len(q) - 1

			// check for same values
			if lastInd != -1 && q[lastInd] == b {
				q = q[:lastInd]
			} else {
				q = append(q, b)
			}
		}

		// store final length
		final[ind] = len(q)
	}

	// get minimum from final slice
	min := final[0]
	for i := 1; i < lr; i++ {
		if v := final[i]; v < min {
			min = v
		}
	}

	return min
}
