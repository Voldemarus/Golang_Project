package main

import "fmt"

func main() {

	sum := 0

	nums := []int{2, 4, 6}

	for _, num := range nums {

		sum += num

	}

	fmt.Println(sum)

}
