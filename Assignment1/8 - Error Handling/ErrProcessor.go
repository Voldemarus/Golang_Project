package main

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	x      float64
	y      float64
	width  float64
	height float64
}

// Implement Error interface to automatically check size
func (r *Rectangle) Error() string {
	if r.width*r.height < 0 {
		return "Invalid dimensions"
	}
	return ""
}

type Interval struct {
	left  float64
	right float64
}

func (interval *Interval) length() (float64, error) {
	if interval.left > interval.right {
		err := errors.New("Interval:: Invalid order of parameters")
		return 0, err
	}
	length := interval.right - interval.left
	return length, nil
}

////////////////////////////

func main() {

	// Built in error processing

	// Generate error in user function
	fmt.Println("\n* Using multiple return values to check for an errors \n")

	varX := 73.0
	varY := 0.0001
	_, err := divideWithCheck(varX, varY)
	if err != nil {
		fmt.Println("divideWithCheck(", varX, ",", varY, ")")
		fmt.Println("cannot perform divideWithCheck() - ", err)
	}

	//  Embedded error checking

	fmt.Println("\n* Error() interface implementation \n")
	a := &Rectangle{12, 40, 10, 10}
	fmt.Println("Rectangle a = ", *a, "err - ", a)
	b := &Rectangle{5, 15, -5, 6}
	fmt.Println("Rectangle b = ", *b, "err - ", b)

	// Error escalation
	fmt.Println("\n* Error status escalation \n")

	aInterval := Interval{5, 10}
	bInterval := Interval{7, 5}

	fmt.Println("aInterval ", aInterval)
	fmt.Println("bInterval ", bInterval)

	resInterval, err := intersect(aInterval, bInterval)
	if err != nil {
		fmt.Println("intersect - ", err)
	} else {
		fmt.Println("Intersection - ", resInterval)
	}

	// Panic and recover
	fmt.Println("\n* Panic/Recover\n")

	fmt.Println("Call panic with recovery activated")
	executePanic(true)

	fmt.Println("Call pnic without recovery")
	executePanic(false)

}

////////////////////////////

func recoveryFunction() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		fmt.Println(recoveryMessage)
	}
	fmt.Println("Recovery function completed\n")
}

func executePanic(callRecovery bool) {
	fmt.Println("executePanic() started")
	if callRecovery {
		defer recoveryFunction() // will be called afterwards in any case!
	}
	panic("Panic event initiated!")
	fmt.Println("executePanic() function finished") // Will not executed!
}

////////////////////////////

func intersect(a, b Interval) (Interval, error) {
	// Check for valid size in the arguments
	_, aerr := a.length()
	if aerr != nil {
		aerr = errors.New("intersect() - substitute old error message")
	}
	_, berr := b.length()
	if berr != nil {
		aerr = fmt.Errorf("intersect()  - add message from embedded error -%v  %v", b, berr)
	}
	resInterval := Interval{0, 0}
	if a.right < b.left || b.right < a.left {
		aerr = errors.New("intersect() - not intersected")
	} else if b.left < a.right {
		resInterval.left = a.right
		resInterval.right = b.left
	} else {
		resInterval.left = b.right
		resInterval.right = a.left
	}

	return resInterval, aerr
}

func divideWithCheck(varX, varY float64) (float64, error) {
	if varY > -0.0002 && varY < 0.0002 {
		err := errors.New("Divisor has too small absolute value")
		return 0, err
	}
	return (varX / varY), nil
}
