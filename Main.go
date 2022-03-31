package main

import (
	"fmt"
)

const (
	_  = iota + 5
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {

	// fmt.Printf("%v, %T\n", i, i)
	// j := strconv.Itoa(i)
	// fmt.Printf("%v, %T\n", j, j)
	// var r rune = 'a'
	// fmt.Printf("%v, %T\n", r, r)
	// var specialistType int
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
}
