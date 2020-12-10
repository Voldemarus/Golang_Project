

From the "Data structures and algorithms in Go", page 67

In permutation method at each recursive call number at index, “i” is swapped with all the numbers that are right of it. Since the number is swapped with all the numbers in its right one by one it will produce all the permutation possible.


Colombina:8-Permutation 1$ go run Pn.go 
Permutation for arrry with length  2
Amount of parmutations -  2
[0 1]
[1 0]


Permutation for arrry with length  3
Amount of parmutations -  6
[0 1 2]
[0 2 1]
[1 0 2]
[1 2 0]
[2 1 0]
[2 0 1]


Permutation for arrry with length  4
Amount of parmutations -  24
[0 1 2 3]
[0 1 3 2]
[0 2 1 3]
[0 2 3 1]
[0 3 2 1]
[0 3 1 2]
[1 0 2 3]
[1 0 3 2]
[1 2 0 3]
[1 2 3 0]
[1 3 2 0]
[1 3 0 2]
[2 1 0 3]
[2 1 3 0]
[2 0 1 3]
[2 0 3 1]
[2 3 0 1]
[2 3 1 0]
[3 1 2 0]
[3 1 0 2]
[3 2 1 0]
[3 2 0 1]
[3 0 2 1]
[3 0 1 2]

