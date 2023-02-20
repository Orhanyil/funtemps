package main

import (
	"flag"
	"fmt"
)

var C float64
var F float64
var K float64
var out string

func init() {

	flag.Float64(&C, "C", 0.0, "temperatur i grader celsius")
	flag.Float64(&F, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64(&K, "K", 0.0, "temperatur i grader kelvin")
	flag.Float64(&out, "out", "C", "valgt resultat: C, F eller K")
}

func main() {
	flag.Parse()
	var res float64 = 0.0

	if out == "C" {
		res = C

	} else if out == "F" {
		res = (C * 9.0 / 5.0) + 32.0

	} else if out == "K" {
		res = C + 273.15
	}

	fmt.Printf("%.2f\n", res)
}
