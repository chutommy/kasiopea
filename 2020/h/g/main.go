package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

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

		// inid grids
		gridA = make([][]string, N)
		gridB = make([][]string, N)
		for n := 0; n < N; n++ {
			gridA[n] = make([]string, M)
			gridB[n] = make([]string, M)
		}

		// reset actions
		actions = [][]int{}

		// get grid A
		for n := 0; n < N; n++ {
			for m := 0; m < M; m++ {
				fmt.Scanf("%1s", &gridA[n][m])
			}
			fmt.Scanf("%1s")
		}
		fmt.Scanf("%1s")

		// get grid B
		for n := 0; n < N; n++ {
			for m := 0; m < M; m++ {
				fmt.Scanf("%1s", &gridB[n][m])
			}
			fmt.Scanf("%1s")
		}
		fmt.Scanf("%1s")

		// solve
		solve(N, M)
		printSolution(f)

		// print status
		fmt.Printf("%d/%d done\n", t+1, T)
	}
}

// L - left
// R - right
// U - up
// D - down
var gridA [][]string
var gridB [][]string
var actions [][]int

func printSolution(f io.Writer) {
	fmt.Fprintln(f, len(actions))
	for _, action := range actions {
		fmt.Fprintf(f, "%d %d\n", action[0], action[1])
	}
}

func solve(N, M int) {

	// check each row
	for n := 0; n < N; n++ {

		// check each piece
		for m := 0; m < M; m++ {

			// rotate if neccessary
			if gridA[n][m] != gridB[n][m] {
				mustRotate(n, m)
			}
		}
	}
}

func mustRotate(y, x int) {

	curr := gridA[y][x]

	// transfer
	if curr == "R" {
		curr = "L"
		x--
	} else if curr == "D" {
		curr = "U"
		y--
	}

	// check if rotable
	for !canRotate(y, x) {
		if curr == "L" {
			// LR
			// ?? case
			mustRotate(y+1, x)

		} else {
			// U?
			// D? case
			mustRotate(y, x+1)
		}
	}

	// rotate
	rotate(y, x)
}

func canRotate(y, x int) bool {

	// LR
	// LR case
	if gridA[y][x] == "L" && gridA[y+1][x] == "L" {
		return true
	}

	// UU
	// DD case
	if gridA[y][x] == "U" && gridA[y][x+1] == "U" {
		return true
	}

	// invalid
	return false
}

func rotate(y, x int) {

	// log the action
	actions = append(actions, []int{y + 1, x + 1})

	if curr := gridA[y][x]; curr == "L" {
		// LR
		// LR case

		gridA[y][x] = "U"
		gridA[y][x+1] = "U"
		gridA[y+1][x] = "D"
		gridA[y+1][x+1] = "D"

	} else {
		// UU
		// DD case

		gridA[y][x] = "L"
		gridA[y][x+1] = "R"
		gridA[y+1][x] = "L"
		gridA[y+1][x+1] = "R"
	}
}
