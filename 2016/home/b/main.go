package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// create output file
	f, err := os.Create("b.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get reader
	r := bufio.NewReader(os.Stdin)

	// get T
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		// get N
		_, _, _ = r.ReadLine()

		// read line
		bs, _, _ := r.ReadLine()

		// solve
		s := solve(bs)
		fmt.Fprintln(f, s)
	}
}

func solve(bs []byte) string {

	// get the R
	key := bs[len(bs)-1]

	// get the move value
	var move int
	for ; move < 26; move++ {

		// get the letter with current move
		ltr := key + byte(move)
		if ltr > 90 {
			ltr -= 26
		}

		// compare
		if 82 == byte(ltr) {
			break
		}
	}

	// get the original message
	orig := make([]byte, len(bs))
	for i, ch := range bs {

		// if space
		if ch == 32 {
			orig[i] = byte(' ')
			continue
		}

		// calculate new letter
		newLtr := ch + byte(move)
		if newLtr > 90 {
			newLtr -= 26
		}

		// store new letter
		orig[i] = newLtr
	}

	return string(orig)
}
