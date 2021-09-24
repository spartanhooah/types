package main

import (
	"fmt"
	"log"
)

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

// structs
type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	// basic types
	myString := "Ben"

	log.Println(myString)

	myString = "John" // new string stored in the variable - strings are immutable

	var myBool = true
	myBool = false
	log.Println(myBool)

	// arrays
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element in array is", myStrings[0])

	// structs
	myCar := Car{
		NumberOfTires: 4,
		Luxury: true,
		BucketSeats: true,
		Make: "Volvo",
		Model: "XC90",
		Year: 2019,
	}

	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
}
