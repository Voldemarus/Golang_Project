package main

import "fmt"

func Permutation(data []int, i int, length int) {
	if length == i {
		// iteration completed, last element in array is reached
		PrintSlice(data)
	} else {
		for j := i; j < length; j++ {
			//replace current symbol with next
			data[i], data[j] = data[j], data[i]
			// calculate permutations for the tail of this state
			Permutation(data, i+1, length)
			// return back initial state
			data[i], data[j] = data[j], data[i]
		}
	}
}

func PrintSlice(data []int) {
	fmt.Println(data)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	arraySize := []int{2, 3, 4}

	for _, size := range arraySize {
		fmt.Println("Permutation for arrry with length ", size)
		fmt.Println("Amount of parmutations - ", factorial(size))
		arr := make([]int, size)
		// fill initial array
		for i := 0; i < size; i++ {
			arr[i] = i
		}
		Permutation(arr[:], 0, len(arr))
		fmt.Println()
		fmt.Println()
	}

}
