package main

import (
	"fmt"
	"math"
)

// Multiple return values, the returned values can be place at the parameter declaration, so that we dont need to declare the variable again.
func calculate(d float64) ( area, circumference float64){
	// Calcualte area
	area = math.Phi * math.Pow(d/2, 2)
	// Calculate circumference
	circumference = math.Phi * d

	// Return two values
	return
}

func main(){
	var area, circumference float64 = calculate(10)

	fmt.Println("Area of circle\t\t:", area)
	fmt.Println("Circumference of circle\t:", circumference)
}