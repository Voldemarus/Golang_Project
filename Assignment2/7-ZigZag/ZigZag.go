package main

import (
	"fmt"
)

type Row struct {
	Data []int
}

type SquareMatrix struct {
	Size int
	Cols []Row
}

func generateSourceMatrix(size int) SquareMatrix {
	if size > 0 && size < 17 {
		newMatrix := SquareMatrix{Size: size}
		matrix := make([]Row, size)
		for i := 0; i < size; i++ {
			rowData := make([]int, size)
			for j := 0; j < size; j++ {
				d := i*size + j // unique index
				rowData[j] = d
			}
			row := Row{rowData}
			matrix[i] = row
		}
		newMatrix.Cols = matrix
		return newMatrix
	}
	return SquareMatrix{Size: 0}
}

func printMatrix(a SquareMatrix) {
	fmt.Printf("      ")
	for i := 0; i < a.Size; i++ {
		fmt.Printf("%2d ", i)
	}
	fmt.Println()
	fmt.Printf("      ")
	for i := 0; i < a.Size; i++ {
		fmt.Printf("---")
	}
	fmt.Println()
	for i := 0; i < a.Size; i++ {
		fmt.Printf(" %2d | ", i)
		r := a.Cols[i]
		fmt.Printf("")
		for j := 0; j < a.Size; j++ {
			fmt.Printf(" %2d", r.Data[j])
		}
		fmt.Println()
	}
}

func (a SquareMatrix) cell(i int, j int) int {
	if i < 0 || j < 0 || i >= a.Size || j >= a.Size {
		return -1 // error indicator
	}
	row := a.Cols[i]
	return row.Data[j]
}

// generates matrix with Hilbert distance as a cell value
func hilbertSpaceMatrix(a SquareMatrix) SquareMatrix {
	if a.Size > 0 && a.Size < 17 {
		newMatrix := SquareMatrix{Size: a.Size}
		matrix := make([]Row, a.Size)
		for i := 0; i < a.Size; i++ {
			rowData := make([]int, a.Size)
			for j := 0; j < a.Size; j++ {
				d := xy2d(a.Size, i, j)
				rowData[j] = d
			}
			row := Row{rowData}
			matrix[i] = row
		}
		newMatrix.Cols = matrix
		return newMatrix
	}
	return SquareMatrix{Size: 0}
}

///// Hilbert curve conversion utilities

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

//Convert  (x,y) to Hilbert distance  d
func xy2d(n int, x int, y int) int {
	d := 0
	for s := n / 2; s > 0; s = s / 2 {
		rx := (x & s) > 0
		ry := (y & s) > 0
		a := 0
		if rx {
			a = 3
		}
		d += s * s * ((a * bool2int(rx)) ^ bool2int(ry))
		x, y = rot(s, x, y, rx, ry)
	}
	return d
}

// rotate/make reflection of quadrant
func rot(n int, x int, y int, rx bool, ry bool) (int, int) {
	if !ry {
		if rx {
			x = n - 1 - x
			y = n - 1 - y
		}
		//Swap x and y
		x, y = y, x
	}
	return x, y
}

///////////////////////

func main() {
	fmt.Println()

	sizeArray := []int{2, 4, 8}
	for _, size := range sizeArray {
		fmt.Printf("\n\nSource Matrix %d*%d\n", size, size)
		a := generateSourceMatrix(size)
		printMatrix(a)
		fmt.Println()
		fmt.Println("Matrix in d space")
		b := hilbertSpaceMatrix(a)
		printMatrix(b)
		fmt.Println()

		linearMatrix := make([]int, size*size)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				index := b.cell(i, j)
				linearMatrix[index] = index
			}
		}
		fmt.Println("Zigzagged representation")
		fmt.Println(linearMatrix)

		fmt.Println()
	}

}
