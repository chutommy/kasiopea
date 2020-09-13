package main

import (
	"fmt"
	"log"
	"os"
)

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

		// get N
		var N int
		fmt.Scan(&N)

		// get volumes
		vols := make([]float32, N)
		for n := 0; n < N; n++ {
			var v float32
			fmt.Scan(&v)
			vols[n] = v
		}

		s1, s2 := solve(vols)
		fmt.Fprintf(f, "%.2f %.2f\n", s1, s2)
	}
}

func solve(vols []float32) (float32, float32) {
	l := len(vols)

	vols = narrow(vols)
	for len(vols) != l {
		l = len(vols)
		vols = narrow(vols)
	}

	volsLen := len(vols)
	// find the max
	var max float32
	for i := 0; i < volsLen; i++ {
		if v := vols[i]; v > max {
			max = v
		}
	}

	fmt.Println(vols)
	return vols[volsLen-1], max
}

func narrow(xi []float32) []float32 {

	var arr []float32
	var prev float32
	var total float32
	var count float32

	for i := 0; i < len(xi); i++ {

		if cur := xi[i]; cur >= prev {

			// add
			total += cur
			prev = cur
			count++

		} else {

			// reset
			arr = append(arr, total/count)
			total = cur
			prev = cur
			count = 1
		}
	}
	arr = append(arr, total/count)

	return arr
}
