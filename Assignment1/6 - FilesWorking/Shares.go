package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type SharePrice struct {
	ticker string
	price  float64 // Current share' price
}

type ShareDividend struct {
	ticker   string
	dividend float64 // dividend in USD per share
	tax      float64 // tax for this particular share
}

var priceList [40]SharePrice
var divList [40]ShareDividend

func main() {
	// Load initial data
	readCSVFile("currentPrice.txt")
	readCSVFile("Dividends.txt")

}

func readCSVFile(fileName string) {
	filePath := "data/" + fileName
	divFile := (fileName == "Dividends.txt")
	csvfile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	index := 0
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Record: ", record)
		if divFile {
			divList[index].ticker = record[0]
			divList[index].dividend = floatValue(record[1])
			divList[index].tax = floatValue(record[2])
			fmt.Println(index, ": dividend", divList[index])
		} else {
			priceList[index].ticker = record[0]
			priceList[index].price = floatValue(record[1])
			fmt.Println(index, ": CurPrice", priceList[index])
		}
		index++
	}
}

func floatValue(input string) float64 {
	if s, err := strconv.ParseFloat(input, 64); err == nil {
		return s
	}
	return 0.0
}
