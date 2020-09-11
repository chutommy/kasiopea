package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// open file for the output
	f, err := os.Create("out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get T
	var T int
	fmt.Scan(&T)

	// range over problem sets
	for t := 0; t < T; t++ {

		// get N, M
		var N, M int
		fmt.Scanf("%d %d", &N, &M)

		houses := make([][]int, N)

		// range over rows
		for n := 0; n < N; n++ {

			// get people
			peopleRow := make([]int, M)

			// range over columns
			for m := 0; m < M; m++ {

				// get K
				var k int
				fmt.Scan(&k)
				peopleRow[m] = k
			}

			houses[n] = peopleRow
		}

		fmt.Fprintln(f, solve(houses))
	}
}

// solve solves the problem from the parsed input
func solve(houses [][]int) int {
	lr := len(houses)
	lc := len(houses[0])

	// count people in rows
	people := make([]int, lr)
	for r := 0; r < lr; r++ {

		sum := 0
		for _, p := range houses[r] {
			sum += p
		}

		people[r] = sum
	}

	// prepare steps grip
	steps := make([][]int, lr)
	for r := 0; r < lr; r++ {
		steps[r] = make([]int, lc)
	}

	// calculate steps to each houses (in rows only)
	for r := 0; r < lr; r++ {
		for c := 0; c < lc; c++ {
			steps[r][c] = solveRow(c, houses[r])
		}
	}

	// calculate current sum of steps for each collumn
	collumnSteps := make([]int, lc)
	for c := 0; c < lc; c++ {
		sum := 0

		for r := 0; r < lr; r++ {
			sum += steps[r][c]
		}
		collumnSteps[c] = sum
	}

	// calculate rowSteps
	rowSteps := make([]int, lr)
	for r := 0; r < lr; r++ {
		rowSteps[r] = solveRow(r, people)
	}

	// find the smallest values
	minCol := collumnSteps[0]
	minRow := rowSteps[0]

	for _, i := range collumnSteps {
		if i < minCol {
			minCol = i
		}
	}
	for _, i := range rowSteps {
		if i < minRow {
			minRow = i
		}
	}

	return minCol + minRow
}

// solveRow calculates the steps if the canteen were in the i-house
func solveRow(i int, row []int) int {
	l := len(row)

	sum := 0
	// add steps
	for n := 0; n < l; n++ {
		sum += abs(i-n) * row[n]
	}
	return sum
}

// abs returns the absolute value of the i
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
