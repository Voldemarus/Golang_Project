
Initial task description : 

https://www.geeksforgeeks.org/print-matrix-zag-zag-fashion/

But to make application more interesting, Hilbert filling curve is generated instead. 

See https://ru.wikipedia.org/wiki/Кривая_Гильберта

See also repository: 

https://github.com/google/hilbert


1. Create square matrix, where each element is calculated as size*Row+Col
2. Estimate distance matrix of the same size, where each element is a distance from the top left corner in Hilbert space
3. Fill linear array (size * size) where each element is placed in accordance to each distance
4. As a result, final linear array is ordered [0,1,2,...]


Example: 

Colombina:7-ZigZag 1$ go run ZigZag.go 



Source Matrix 2*2
       0  1 
      ------
  0 |   0  1
  1 |   2  3

Matrix in d space
       0  1 
      ------
  0 |   0  1
  1 |   3  2

Zigzagged representation
[0 1 2 3]



Source Matrix 4*4
       0  1  2  3 
      ------------
  0 |   0  1  2  3
  1 |   4  5  6  7
  2 |   8  9 10 11
  3 |  12 13 14 15

Matrix in d space
       0  1  2  3 
      ------------
  0 |   0  3  4  5
  1 |   1  2  7  6
  2 |  14 13  8  9
  3 |  15 12 11 10

Zigzagged representation
[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]



Source Matrix 8*8
       0  1  2  3  4  5  6  7 
      ------------------------
  0 |   0  1  2  3  4  5  6  7
  1 |   8  9 10 11 12 13 14 15
  2 |  16 17 18 19 20 21 22 23
  3 |  24 25 26 27 28 29 30 31
  4 |  32 33 34 35 36 37 38 39
  5 |  40 41 42 43 44 45 46 47
  6 |  48 49 50 51 52 53 54 55
  7 |  56 57 58 59 60 61 62 63

Matrix in d space
       0  1  2  3  4  5  6  7 
      ------------------------
  0 |   0  1 14 15 16 19 20 21
  1 |   3  2 13 12 17 18 23 22
  2 |   4  7  8 11 30 29 24 25
  3 |   5  6  9 10 31 28 27 26
  4 |  58 57 54 53 32 35 36 37
  5 |  59 56 55 52 33 34 39 38
  6 |  60 61 50 51 46 45 40 41
  7 |  63 62 49 48 47 44 43 42

Zigzagged representation
[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63]

Colombina:7-ZigZag 1$ 