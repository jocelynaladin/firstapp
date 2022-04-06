package main

import "fmt"

// const (
// 	_  = iota + 5
// 	KB = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// 	PB
// )

// type Doctor struct {
// 	number     int
// 	actorName  string
// 	companions []string
// }

type Animal struct {
	Name   string
	origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// fileSize := 4000000000.
	// fmt.Printf("%.2fGB\n", fileSize/GB)
	// aDoctor := Doctor{
	// 	number:    3,
	// 	actorName: "Jocelyn Aladin",
	// 	companions: []string{
	// 		"Nadine Aladin",
	// 		"Philippe Henry",
	// 		"Natt Aladin",
	// 	},
	// }
	// fmt.Println(aDoctor)
	b := Bird{}
	b.Name = "Emu"
	b.origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false

	fmt.Println(b)
}
