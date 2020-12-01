
Quick sort implementation


Based on 
"DATA STRUCTURES & ALGORITHMS IN GO" FIRST EDITION HEMANT JAIN

Note! Implementation in the boo contains bugs!!!

Constraints:
- single thread
- Globally defined array to hold original/sorted data


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

go run QuickSort.go 

Example: 

Source array
[{0 19 Bob} {0 18 Albert} {0 47 Victor} {0 8 Albert} {1 5 Tamara} {1 36 Elizabeth} {0 21 Jake} {0 2 Mike} {1 26 Elizabeth} {1 1 Ann} {1 38 Barbara} {1 3 Lily} {1 2 Jane} {1 24 Nataly} {0 25 Stephan} {0 32 Mike} {0 36 Frank} {0 14 Tom} {1 12 Jane} {1 48 Barbara} {0 31 Jake} {0 41 Jake} {1 21 Ann} {1 47 Juliet} {0 13 Paul} {0 5 Stephan} {0 0 Joe} {0 22 Mike} {0 12 Mike} {0 24 Tom} {1 15 Tamara} {1 44 Nataly} {0 39 Bob} {1 45 Tamara} {0 7 Victor} {0 23 Paul} {1 7 Juliet} {1 27 Juliet} {0 4 Tom} {1 0 Mary} {0 6 Frank} {1 17 Juliet} {1 8 Barbara} {1 11 Ann} {0 16 Frank} {1 42 Jane} {1 43 Lily} {0 9 Bob} {1 30 Mary} {0 43 Paul} {0 10 Joe} {1 28 Barbara} {0 48 Albert} {1 10 Mary} {0 37 Victor} {1 22 Jane} {0 26 Frank} {0 40 Joe} {1 25 Tamara} {0 42 Mike} {1 37 Juliet} {0 38 Albert} {1 35 Tamara} {0 49 Bob} {1 41 Ann} {0 45 Stephan} {1 16 Elizabeth} {1 18 Barbara} {1 40 Mary} {1 34 Nataly} {1 33 Lily} {1 46 Elizabeth} {0 3 Paul} {1 31 Ann} {0 17 Victor} {1 4 Nataly} {0 30 Joe} {1 19 Helga} {1 14 Nataly} {1 32 Jane} {0 35 Stephan} {0 34 Tom} {0 11 Jake} {1 6 Elizabeth} {1 39 Helga} {0 46 Frank} {1 13 Lily} {1 29 Helga} {0 44 Tom} {0 29 Bob} {1 9 Helga} {0 20 Joe} {0 15 Stephan} {0 28 Albert} {0 33 Paul} {1 20 Mary} {0 1 Jake} {0 27 Victor} {1 23 Lily} {1 49 Helga}]

Sorted array
[{0 0 Joe} {1 0 Mary} {0 1 Jake} {1 1 Ann} {0 2 Mike} {1 2 Jane} {0 3 Paul} {1 3 Lily} {0 4 Tom} {1 4 Nataly} {0 5 Stephan} {1 5 Tamara} {0 6 Frank} {1 6 Elizabeth} {0 7 Victor} {1 7 Juliet} {0 8 Albert} {1 8 Barbara} {0 9 Bob} {1 9 Helga} {0 10 Joe} {1 10 Mary} {0 11 Jake} {1 11 Ann} {0 12 Mike} {1 12 Jane} {0 13 Paul} {1 13 Lily} {0 14 Tom} {1 14 Nataly} {0 15 Stephan} {1 15 Tamara} {0 16 Frank} {1 16 Elizabeth} {0 17 Victor} {1 17 Juliet} {0 18 Albert} {1 18 Barbara} {0 19 Bob} {1 19 Helga} {0 20 Joe} {1 20 Mary} {0 21 Jake} {1 21 Ann} {0 22 Mike} {1 22 Jane} {0 23 Paul} {1 23 Lily} {0 24 Tom} {1 24 Nataly} {0 25 Stephan} {1 25 Tamara} {0 26 Frank} {1 26 Elizabeth} {0 27 Victor} {1 27 Juliet} {0 28 Albert} {1 28 Barbara} {0 29 Bob} {1 29 Helga} {0 30 Joe} {1 30 Mary} {0 31 Jake} {1 31 Ann} {0 32 Mike} {1 32 Jane} {0 33 Paul} {1 33 Lily} {0 34 Tom} {1 34 Nataly} {0 35 Stephan} {1 35 Tamara} {0 36 Frank} {1 36 Elizabeth} {0 37 Victor} {1 37 Juliet} {0 38 Albert} {1 38 Barbara} {0 39 Bob} {1 39 Helga} {0 40 Joe} {1 40 Mary} {0 41 Jake} {1 41 Ann} {0 42 Mike} {1 42 Jane} {0 43 Paul} {1 43 Lily} {0 44 Tom} {1 44 Nataly} {0 45 Stephan} {1 45 Tamara} {0 46 Frank} {1 46 Elizabeth} {0 47 Victor} {1 47 Juliet} {0 48 Albert} {1 48 Barbara} {0 49 Bob} {1 49 Helga}]
Colombina:1-Quicksort 1$ 