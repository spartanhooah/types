package main

import "log"

// basic types (numbers, strings, booleans)
var myInt int

// Discouraged
var myInt16 int16
var myInt32 int32
var myInt64 int64

// unsigned int
var myUint uint

// floats/decimals
var myFloat32 float32
var myFloat64 float64

func main() {
	myString := "Ben"

	log.Println(myString)

	myString = "John" // new string stored in the variable - strings are immutable

	var myBool = true
	myBool = false
	log.Println(myBool)
}