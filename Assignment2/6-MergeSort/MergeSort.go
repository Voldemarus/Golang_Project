package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
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

func (s Spouse) print() {
	g := "Female"
	if s.gender == GenderMale {
		g = "Male"
	}
	fmt.Println("fmaily #", s.num, " Name:", s.name, " Gender:", g)
}

const (
	amountOfSpouses = 200 // total amount of spouses in the array
)

func generateSpouseRandomList() []Spouse {
	index := 0
	var spouseArr []Spouse
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
	return spouseArr
}

//
// Create CSV file with unsorted data
//
func createCSVFile(spouseArr []Spouse, fileName string) error {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	defer w.Flush()
	for _, spEntry := range spouseArr {
		var record []string
		record = append(record, spEntry.name)
		record = append(record, fmt.Sprintf("%d", spEntry.gender))
		record = append(record, fmt.Sprintf("%d", spEntry.num))

		w.Write(record)
		if err := w.Error(); err != nil {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////

func main() {
	// Create initial array of structures, which will be sorted
	spouseArr := generateSpouseRandomList()
	// Save data into file to implement external merge sorting
	err := createCSVFile(spouseArr, "unsortedData.csv")
	if err != nil {
		log.Fatalln("Error during CSV creation -", err)
	}
	// Now we will sort data and write sorted data into another CSV file
	err = mergeSort("unsortedData.csv", "sortedCSV.txt")

	fmt.Println("MergeSort = Initial unsorted set written to CSV file")
}

////////////////////////   Ectetrnal Merge Sort ///////////////////////

//
// External merge sort
//
func mergeSort(inputFile, outputFile string) error {

	return nil
}
