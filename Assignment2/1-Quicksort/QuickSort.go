package main

import (
	"fmt"
	"math/rand"
)

const (
	GenderMale = iota
	GenderFemale
)

type Spouse struct {
	gender int // gender identifer
	num    int // family identifier
	name   string
}

type SpouseUID interface {
	id() int // unified identifier
}

func (s Spouse) id() int {
	return s.num*2*amountOfSpouses + s.gender
}

const (
	amountOfSpouses = 50 // total amount of spouses in the array
)

var spouseArr []Spouse

func generateSpouseRandomList() {
	index := 0
	maleNames := [...]string{"Joe", "Jake", "Mike", "Paul", "Tom", "Stephan", "Frank",
		"Victor", "Albert", "Bob"}

	femaleNames := [...]string{"Mary", "Ann", "Jane", "Lily", "Nataly",
		"Tamara", "Elizabeth", "Juliet", "Barbara", "Helga"}

	// We'll create males part first and put them into array
	for i := 0; i < amountOfSpouses; i++ {
		sp := Spouse{gender: GenderMale, num: i, name: maleNames[index]}
		spouseArr = append(spouseArr, sp)
		index++
		if index >= len(maleNames) {
			index = 0
		}
	}
	// Now add female part to the same list
	index = 0
	for i := 0; i < amountOfSpouses; i++ {
		sp := Spouse{gender: GenderFemale, num: i, name: femaleNames[index]}
		spouseArr = append(spouseArr, sp)
		index++
		if index >= len(femaleNames) {
			index = 0
		}
	}
	//
	// Now we'll make random permutation of entires in array
	//
	for i := 0; i < amountOfSpouses*15; i++ {
		swapIndLeft := rand.Intn(amountOfSpouses*2 - 1)
		swapIndRight := rand.Intn(amountOfSpouses*2 - 1)

		spouseArr[swapIndLeft], spouseArr[swapIndRight] =
			spouseArr[swapIndRight], spouseArr[swapIndLeft]
	}
}

func main() {
	// Create initial array of structures, which will be sorted
	generateSpouseRandomList()

	fmt.Println("Source array")
	fmt.Println(spouseArr)

	quickSortArray()
	fmt.Println("Sorted array")
	fmt.Println(spouseArr)

}

////////////////////////   Quick Sorting ///////////////////////

func quickSortArray() {

	upper := len(spouseArr) - 1
	quickSort1(0, upper)
}

func quickSort1(lower int, upper int) {
	// if both indexes are met, or overlapped then return
	if upper <= lower {
		return
	}
	// Pick up initial point, as a first element of array
	seed := spouseArr[lower]
	// set up indexes
	start := lower
	end := upper

	// Now check for consistency and update indexes
	for lower < upper {
		if spouseArr[lower].id() < seed.id() && lower < upper {
			// current element is less then seed, so it is on proper place and we
			// can move left index up
			lower++
		} else if spouseArr[upper].id() > seed.id() && lower <= upper {
			// element on the right end is greater then seed, so we do not need to move it
			// Use less restrictive comparasion to cover all cases
			upper--
		} else if lower < upper {
			// now make swap for  elements with updated indexes
			spouseArr[lower], spouseArr[upper] = spouseArr[upper], spouseArr[lower]
		}
	}
	// move seed value back to start position
	//
	// Init recursive calls. Isolate central point and
	// make sorting for left and right parts
	//
	quickSort1(start, upper-1)
	quickSort1(upper+1, end)

}
