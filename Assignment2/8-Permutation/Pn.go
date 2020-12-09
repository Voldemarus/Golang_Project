package main

import "fmt"

type Vector struct {
	Data []int
}

type Permutation struct {
	Vectors []Vector
}

// Returns size of the vecotrs, stored in Permutation list
func (p Permutation) size() int {
	if len(p.Vectors) > 0 {
		return len(p.Vectors[0].Data)
	}
	return 0
}

func (p Permutation) item(index int) Vector {
	if index < 0 || index > (p.size()-1) {
		v := new(Vector)
		return *v // create and return empty vector
	}
	return p.Vectors[index]
}

func (p Permutation) printAllItems() {
	for i, v := range p.Vectors {
		fmt.Printf("%2d : %v\n", i, v)
	}
	fmt.Println()
}

func PermutationCreate(arr []int) *Permutation {
	len := len(arr)
	if len <= 0 || len > 10 {
		return nil
	}
	p := new(Permutation)

	return p
}

func main() {
	arraySize := []int{3, 4, 6, 8}

	for _, size := range arraySize {
		fmt.Println("Permutation for arrry with length ", size)
		arr := make([]int, size)
		// fill initial array
		for i := 0; i < size; i++ {
			arr[i] = i
		}

		p := PermutationCreate(arr)
		p.printAllItems()

		fmt.Println()
	}

}
