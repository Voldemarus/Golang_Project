
Merge sort implementation


Based on 
"DATA STRUCTURES & ALGORITHMS IN GO" FIRST EDITION HEMANT JAIN


There is an array of Spouse records, which contains gender, name and identification number of family
Records are ordered randomly. We hould align them by family id in a such way, that male is the first and the female is the second.

Structure to Sort:

type Spouse struct {
	gender int // gender identifer : 0- Male, 1 - Female
	num    int // family identifier
	name   string
}

Comparsion Rule:

func (s Spouse) id() int {
	return s.num*2*amountOfSpouses + s.gender
}

To run program:

go run MergeSort.go 

Example: 

go run MergeSort.go 

Initial chunk
 1 40 30 11 21 20 10 31  0 41 

Sorted chunk
 0  1 10 11 20 21 30 31 40 41 

 ---- 

 Initial unsorted data are placed into unSortedData.csv
 Sorted data are placed into sortedCSV.txt (txt! to simplify view on Mac platform)

 ----

Colombina:6-MergeSort 1$ cat sortedCSV.txt 
Joe,0,0
Mary,1,0
Jake,0,1
Ann,1,1
Mike,0,2
Jane,1,2
Paul,0,3
Lily,1,3
Tom,0,4
Nataly,1,4
