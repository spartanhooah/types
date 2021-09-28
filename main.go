package main

import (
	"fmt"
	"sort"
	// "github.com/eiannone/keyboard"
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

// functions
type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (animal *Animal) Says() {
	fmt.Printf("A %s says %s\n", animal.Name, animal.Sound)
}

func (animal *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs\n", animal.Name, animal.NumberOfLegs)
}

// channels
var keyPressChan chan rune

// interfaces
// interface type
type Pet interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (dog *Dog) Says() string {
	return dog.Sound
}

func (dog *Dog) HowManyLegs() int {
	return dog.NumberOfLegs
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (cat *Cat) Says() string {
	return cat.Sound
}

func (cat *Cat) HowManyLegs() int {
	return cat.NumberOfLegs
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
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
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

	//maps
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	delete(intMap, "four")

	element, found := intMap["four"]

	if found {
		fmt.Println("Found", element)
	} else {
		fmt.Println("Didn't find", element)
	}

	intMap["two"] = 4

	// functions
	myTotal := sumMany(2, 3, 4, 5)

	fmt.Println(myTotal)

	dog := Animal{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}

	cat.Says()
	cat.HowManyLegs()

	// channels
	// keyPressChan = make(chan rune)
	// go listenForKeyPress()

	// fmt.Println("Press any key, or q to quit")
	// _ = keyboard.Open()

	// defer func() {
	// 	keyboard.Close()
	// }()

	// for {
	// 	char, _, _ := keyboard.GetSingleKey()

	// 	if char == 'q' || char == 'Q' {
	// 		break
	// 	}

	// 	keyPressChan <- char
	// }

	// interfaces
	dog2 := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	riddle(&dog2)

	cat2 := Cat{
		Name:         "cat",
		NumberOfLegs: 4,
		Sound:        "meow",
		HasTail:      true,
	}

	riddle(&cat2)

	// expressions
	age := 10
	name := "Jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Right handed: %t\n", name, age, rightHanded)

	ageInTenYears := age + 10
	fmt.Printf("In ten years, %s will be %d years old\n", name, ageInTenYears)

	isATeenager := age >= 13 && age < 20
	fmt.Println(name, "is a teenager:", isATeenager)
}

func changePointerValue(num *int) {
	*num = 25
}

func deleteFromSlice(source []string, i int) []string {
	source[i] = source[len(source)-1]
	// source[len(source) - 1] = ""

	return source[:len(source)-1]
}

// "naked" return; generally discouraged
func addtwoNumbers(x, y int) (sum int) {
	sum = x + y

	return
}

func sumMany(nums ...int) int {
	total := 0

	for _, addend := range nums {
		total += addend
	}

	return total
}

// channels
func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}

// interfaces
func riddle(pet Pet) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, pet.Says(), pet.HowManyLegs())
	fmt.Println(riddle)
}
