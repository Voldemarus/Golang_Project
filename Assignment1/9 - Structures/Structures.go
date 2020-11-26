package main

import (
	"fmt"
	"reflect"
)

//
// Define structures with basic fields
//

type Figure interface {
	area() float64 // get area of figure
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	diameter float64
}

//
// Defne methods for particular types, to calculate derived (transit) field - area
//
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Circle) area() float64 {
	return r.diameter * r.diameter * 3.1415926 / 4.0
}

func equalTo(a, b Figure) bool {
	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(b)
	if t1 != t2 {
		return false
	}
	return a.area() == b.area()
}

//
// Now define interface based method
//
func biggerFigure(a, b Figure) Figure {
	if a.area() > b.area() {
		return a
	} else {
		return b
	}
}

func main() {

	rectangle := Rectangle{width: 10, height: 10}
	circle := Circle{5.0}

	fmt.Println("Rectangle area ", rectangle.area())
	fmt.Println("Circle area ", circle.area())

	bFigure := biggerFigure(rectangle, circle)
	s := fmt.Sprintf("Bigger is the %T", bFigure)
	fmt.Println(s)

	m := map[string]Figure{
		"Figure1": Rectangle{width: 20, height: 20},
		"Figure2": Circle{15},
		"Figure3": Circle{17},
		"Figure4": Rectangle{30, 4},
		"Figure5": Rectangle{25, 13},
	}

	biggestFigure := Figure(Rectangle{0, 0})
	biggestID := "Not found"
	for _, value := range m {
		biggestFigure = biggerFigure(biggestFigure, value)
	}

	fmt.Println("Biggest figure area is ", biggestFigure.area())

	// As GO cannot search maps by value (as Obj-C, for example) we need to scan agai
	// and define additional methods for interface!

	for k, v := range m {
		if equalTo(v, biggestFigure) {
			biggestID = k
			break
		}
	}

	fmt.Println("Its name - ", biggestID)
}
