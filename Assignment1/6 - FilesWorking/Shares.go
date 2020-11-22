package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
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
	period   float64 // amount of months between dividend payments
}

type ShareIncome struct {
	ticker    string
	income    float64 // additional income for current investment (avg. month)
	numShares float64 // amount of shares
}

type IncomeList []ShareIncome

// func (a IncomeList) Len() int           { return len(a) }
// func (a IncomeList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a IncomeList) Less(i, j int) bool { return a[i].income < a[j].income }

const (
	kPortfolioSize      int     = 40   // max amount of shares in portfolio
	kBetterSharesAmount int     = 3    // amount of shares placed into report
	kConsolePrint       bool    = true // duplicate report to console
	kBbrokerFee         float64 = 10.0 // Broker's fee added to any purchase/sell opperation
)

var priceList [kPortfolioSize]SharePrice
var divList [kPortfolioSize]ShareDividend
var income [kPortfolioSize]ShareIncome
var currentPortfolioSize = kPortfolioSize

var investingSum float64 = 1000.0 // By default consider we will invest this sum

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 {
		investingSum = floatValue(argsWithoutProg[0])
	}

	// Load initial data
	readCSVFile("currentPrice.txt")
	readCSVFile("Dividends.txt")

	// Process data and fill income array
	calculateIncome()
	var incomeData IncomeList = income[0:currentPortfolioSize]

	// sort slice (and underlying array as well) to liit sorting procedure
	// with non-empty records
	sort.Slice(incomeData, func(i, j int) bool {
		return incomeData[i].income > incomeData[j].income
	})

	// and generate output file with suggestions
	prepareReport(kConsolePrint, kBetterSharesAmount)

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
		// fmt.Println("Record: ", record)
		if divFile {
			divList[index].ticker = record[0]
			divList[index].dividend = floatValue(record[1])
			divList[index].tax = floatValue(record[2])
			divList[index].period = floatValue(record[3])
			//			fmt.Println(index, ": dividend", divList[index])
		} else {
			priceList[index].ticker = record[0]
			priceList[index].price = floatValue(record[1])
			//			fmt.Println(index, ": CurPrice", priceList[index])
		}
		index++
	}
}

func calculateIncome() {
	currentPortfolioSize = 0
	for index := range priceList {
		cPrice := priceList[index] // current price for item i
		cDivPar := divList[index]  // dividend parameters for the same item
		if len(cPrice.ticker) > 0 {
			// Step 1. Calculate amount of shares we can buy with
			// current price
			amountToBuy := math.Floor((investingSum - kBbrokerFee) / cPrice.price)
			// Step 2. Calculate dividend, which will be paid each month (in average)
			dividendBrutto := amountToBuy * cDivPar.dividend
			dividendNetto := dividendBrutto * (1.0 - cDivPar.tax)
			dividendMonth := dividendNetto / cDivPar.period
			//	 fmt.Println("ticker - ", cPrice.ticker, " can buy -", amountToBuy,
			//				" per Month -", dividendMonth)
			income[index].ticker = cPrice.ticker
			income[index].income = dividendMonth
			income[index].numShares = amountToBuy
			currentPortfolioSize++
		}

	}
}

//
//  consolePrint bool - duplicate output to console if true
//  numShare int - amount of shares shown in the report
//
func prepareReport(consolePrint bool, numShare int) {

	outputData := income[0:numShare]

	outFile, err := os.Create("report.txt")
	if err != nil {
		log.Fatalln("Couldn't create  output file", err)
		return
	}
	defer outFile.Close()

	reportHeader := fmt.Sprintf("Invested sum: %7.2f\n\n", investingSum)
	reportHeader = reportHeader + "Ticker\t\tAmount of shares\tPlanned monthly income\n"
	_, err = outFile.WriteString(reportHeader)
	if err != nil {
		log.Fatalln("Error during writing to output file", err)
		return
	}
	for index := 0; index < numShare; index++ {
		currentIncome := income[index]
		tickerLine := currentIncome.ticker + "\t\t\t" + fmt.Sprintf("%.0f", currentIncome.numShares) +
			"\t\t\t" + fmt.Sprintf("%.2f YSD/month\n", currentIncome.income)
		_, err := outFile.WriteString(tickerLine)
		if err != nil {
			log.Fatalln("Error during writing to output file", err)
			return
		}
	}

	fmt.Println(outputData)
}

////////// Service functions ///////////////

func floatValue(input string) float64 {
	if s, err := strconv.ParseFloat(input, 64); err == nil {
		return s
	}
	return 0.0
}
