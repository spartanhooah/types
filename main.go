package main

import (
	"fmt"
	"sort"
)

// basic types (numbers, strings, booleans)
var myInt int

// (discouraged)
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

	fmt.Println(myString)

	myString = "John" // new string stored in the variable - strings are immutable

	var myBool = true
	myBool = false
	fmt.Println(myBool)

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

	fmt.Printf("My car is a %d %s %s\n", myCar.Year, myCar.Make, myCar.Model)

	// reference types
	// pointers
	x := 10
	myFirstPointer := &x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15

	fmt.Println("x is now", x)
	
	changePointerValue(myFirstPointer)
	
	fmt.Println("x is now", x)

	// slices
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, animal := range animals {
		fmt.Println(i, animal)
	}

	fmt.Println("Element 0 is", animals[0])
	fmt.Println("The first two elements are", animals[0:2])
	fmt.Println("The slice has", len(animals), "animals")
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = deleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func changePointerValue(num *int) {
	*num = 25
}

func deleteFromSlice(source []string, i int) []string {
	source[i] = source[len(source) - 1]
	// source[len(source) - 1] = ""
	
	return source[:len(source) - 1]
}
