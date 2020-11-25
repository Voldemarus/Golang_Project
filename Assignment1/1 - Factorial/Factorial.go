package main

import (
	"fmt"
)

func main() {
	sumRes := 0
	sum(1, 5, &sumRes)

	fmt.Println("sum = ", sumRes)

	s, g := factorial(7)
	fmt.Println("Factorial ", s, "= ", g)

}

func factorial(m int) (int, int) {
	result := 1
	if m > 1 {
		_, result := factorial(m - 1)
		return m, result * m
	}
	return m, result
}

func sum(from int, to int, res *int) {
	fmt.Println("from = ", from, "to = ", to)
	if from >= to {
		*res = from
		return
	} else if (to - from) == 1 {
		*res = from + to
		return
	}
	middle := (to + from) / 2
	middle2 := middle + 1
	if middle2 == from {
		middle2++
	}
	var left, right int

	sum(from, middle, &left)
	sum(middle2, to, &right)

	*res = left + right
}
