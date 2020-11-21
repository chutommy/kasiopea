package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// L - left
// R - right
// U - up
// D - down
var gridA [][]string
var gridB [][]string
var actions [][]int

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

		// init actions
		actions = [][]int{}

		// get grid A
		for n := 0; n < N; n++ {
			for m := 0; m < M; m++ {
				// store piece
				var p string
				fmt.Scanf("%1s", &p)
				gridA[n][m] = p
			}
			fmt.Scanf("%1s")
		}

		fmt.Scanf("%1s")

		// get grid B
		for n := 0; n < N; n++ {
			for m := 0; m < M; m++ {
				// store piece
				var p string
				fmt.Scanf("%1s", &p)
				gridB[n][m] = p
			}
			fmt.Scanf("%1s")
		}

		fmt.Scanf("%1s")

		// solve
		solve(N, M)
		printSolution(f)
		fmt.Printf("%d/%d done\n", t+1, T)
	}
}

func printGrid() {
	for _, row := range gridA {
		fmt.Println(row)
	}
}

func printActions() {
	for _, action := range actions {
		fmt.Println(action)
	}
}

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

			//TODO remove this code
			if gridA[n][m] != gridB[n][m] {
				panic("WTF 00")
			}
			// to check invalid situation
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

	//TODO remove this code
	if !canRotate(y, x) {
		panic("WTF 01")
	}
	// to check invalid situation

	// rotate
	rotate(y, x)
}

func canRotate(y, x int) bool {

	// L case
	if gridA[y][x] == "L" && gridA[y+1][x] == "L" {
		return true

	}

	// U case
	if gridA[y][x] == "U" && gridA[y][x+1] == "U" {
		return true

	}

	// invalid
	return false
}

func rotate(y, x int) {

	// log the action
	actions = append(actions, []int{y + 1, x + 1})

	curr := gridA[y][x]

	// LR
	// LR case
	if curr == "L" {

		//TODO remove this code
		if gridA[y][x+1] != "R" || gridA[y+1][x] != "L" {
			panic("WTF 02")
		}
		// to check invalid situation

		gridA[y][x] = "U"
		gridA[y][x+1] = "U"
		gridA[y+1][x] = "D"
		gridA[y+1][x+1] = "D"

		// UU
		// DD case
	} else {

		//TODO remove this code
		if gridA[y+1][x] != "D" || gridA[y][x+1] != "U" {
			panic("WTF 03")
		}
		// to check invalid situation

		gridA[y][x] = "L"
		gridA[y][x+1] = "R"
		gridA[y+1][x] = "L"
		gridA[y+1][x+1] = "R"

	}
}
