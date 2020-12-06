package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

const (
	GenderMale = iota
	GenderFemale
)

const (
	amountOfSpouses = 5 // total amount of spouses in the array
	kChunkSize      = 5 // amount of record in one chunk
	kThreadsAmount  = 6 // amount of patallel threads to process data
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

func (s Spouse) print(full bool) {
	g := "Female"
	if s.gender == GenderMale {
		g = "Male"
	}
	if full {
		fmt.Println("Family #", s.id(), "Num: ", s.num, " Name:", s.name, " Gender:", g)
	} else {
		fmt.Println("#", s.id())
	}
}

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
// Create CSV file from slice
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

////////////////////////   Ectetrnal Merge Sort ///////////////////////

//
// name - file name,
// offset - amount of records from the beginning of the file to be skipped
// recordCount - amount of records (chunk size) or -1 to
// read the whole file at once
//
func readDatafromCSVFile(csvfile *os.File, recordCount int) ([]Spouse, bool) {

	// Parse the file
	r := csv.NewReader(csvfile)
	finished := false
	var result []Spouse
	index := 0
	for {
		record, err := r.Read()
		if err == io.EOF || (recordCount > 0 && index == recordCount) {
			// We should break reading loop if EOF is reached
			// or predefined amount of records are loaded
			finished = true
			break
		}
		if index >= 0 {
			aName := record[0]
			aGender, _ := strconv.Atoi(record[1])
			aNum, _ := strconv.Atoi(record[2])
			newSpouse := Spouse{gender: aGender, num: aNum, name: aName}
			result = append(result, newSpouse)
		}
	}
	return result, finished
}

//
// External merge sort
//
func mergeSort(inputFile, outputFile string) error {
	// Process input file chubk by chunk
	csvfile, err := os.Open(inputFile)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()

	// Parse the file
	topWg := new(sync.WaitGroup)
	finished := false
	chunks := 0 // total amount of chunks
	var chunkToSort []Spouse
	for finished == false {
		// get chunk
		chunkToSort, finished = readDatafromCSVFile(csvfile, kChunkSize)
		if len(chunkToSort) > 0 {
			topWg.Add(1)
			go processMergeSorting(chunkToSort, chunks, topWg) // merge sorting for chunk and storing it in the file
			chunks++                                           // increment counter
		}
	}
	topWg.Wait()
	fmt.Println("Chunks processed. Total ", chunks, " files created")

	// Now we should merge separate chunk files

	fmt.Println("Final merge from sorted chunks")
	if chunks == 1 {
		// only one chunk created, so we can just move it into output
		// filename
		chunkName := "tmp/chunk0.csv"
		return os.Rename(chunkName, outputFile)
	} else {
		// complex case, we need to process severlchunks and make united
		// output file
	}
	return nil
}

// perform in memory merge sorting for chunk and store it in the temporary file

func processMergeSorting(chunk []Spouse, chunkNum int, wg *sync.WaitGroup) {
	//	size := len(chunk) // can be less then defined in const!

	fmt.Println("Initial chunk")
	spouseListPrint(chunk, false)
	fmt.Println()
	output := mergeSorting(chunk, nil) // call merge processing on the root node
	fmt.Println("Sorted chunk")
	spouseListPrint(output, false)
	fmt.Println()

	// Now we have sorted chunk, store it into temporary file
	fileName := fmt.Sprintf("tmp/chunk%d.csv", chunkNum)
	err := createCSVFile(output, fileName)
	if err != nil {
		log.Fatalln("Error on chnk file creation -", err)
	}
	defer wg.Done()
}

func mergeSorting(input []Spouse, wg *sync.WaitGroup) []Spouse {

	if len(input) <= 1 {
		return input // end of recursion reached, no changes required
	}
	//	defer wg.Done()

	pivot := len(input) / 2

	localWg := new(sync.WaitGroup)
	//	localWg.Add(2)
	left := input[:pivot]
	right := input[pivot:]

	left = mergeSorting(left, localWg)
	right = mergeSorting(right, localWg)

	output := mergeLR(left, right)

	return output
}

func mergeLR(left []Spouse, right []Spouse) []Spouse {
	leftStop := len(left) - 1
	if leftStop < 0 {
		return right
	}
	rightStop := len(right) - 1
	if rightStop < 0 {
		return left
	}

	out := make([]Spouse, len(left)+len(right))
	leftIndex, rightIndex, index := 0, 0, 0
	for leftIndex <= leftStop && rightIndex <= rightStop {
		if left[leftIndex].id() < right[rightIndex].id() {
			// data on the left is less then on the right side,
			// so we put into tmp array element from the left half
			out[index] = left[leftIndex]
			leftIndex++
		} else {
			// cover both == and > cases
			// value on the right should be placed to output position
			out[index] = right[rightIndex]
			rightIndex++
		}
		index++
	} // for
	for leftIndex <= leftStop {
		out[index] = left[leftIndex]
		index++
		leftIndex++
	}
	for rightIndex <= rightStop {
		out[index] = right[rightIndex]
		index++
		rightIndex++
	}

	return out
}

//
// Main
//

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

/// Auxillary debug functions ///

func compactSpousePrint(arr []Spouse, from int, to int) {
	if from < 0 {
		from = 0
		to = len(arr) - 1
	}
	for i := from; i <= to; i++ {
		fmt.Printf("%2d ", arr[i].id())
	}
	fmt.Println()
}

func spouseListPrint(arr []Spouse, full bool) {
	if full {
		for i, v := range arr {
			fmt.Printf("%2d :: ", i)
			if len(v.name) > 0 {
				v.print(full)
			} else {
				fmt.Printf("--")
			}
		}
	} else {
		for _, v := range arr {
			fmt.Printf("%2d ", v.id())
		}
		fmt.Println()
	}

}
