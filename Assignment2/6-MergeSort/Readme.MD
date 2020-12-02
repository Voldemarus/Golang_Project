
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

