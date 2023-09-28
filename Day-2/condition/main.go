/* Conditional Statements */
/*
	1. if
	2. else if
	3. else
*/

package main

import (
	"fmt"
)

func main() {
	// Example 

	var point int = 8

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Printf("Tidak lulus. Nilai anda %d\n", point)
	}

	// Temporary variable in if statement

	var point2 float32 = 8840.0

	if percent:=point2/100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// Switch Case

	var point3 int = 6

	switch point3 {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

	// Switch Case with multiple condition

	var point4 int = 2

	switch point4 {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")
	default:
		// can be use with {} to make it more readable
		{
			fmt.Println("Not Bad")
			fmt.Println("You can be better!")
		}

	}

	// Switch Case with if else style
	
	var point5 int = 6

	switch {
	case point5 == 8:
		fmt.Println("Perfect")
	case (point5 < 8) && (point5 > 3):
		fmt.Println("Awesome")
	default:
		// can be use with {} to make it more readable
		{
			fmt.Println("Not Bad")
			fmt.Println("You can be better!")
		}
	}

	// Fallthrough

	var point6 int = 4

	switch {
	case point6 == 8:
		fmt.Println("Perfect")
	case (point6 < 8) && (point6 > 3):
		fmt.Println("Awesome")
		fallthrough
	case point6 < 5:
		fmt.Println("You need to learn more")
	default:
		// can be use with {} to make it more readable
		{
			fmt.Println("Not Bad")
			fmt.Println("You can be better!")
		}
	}

	//Nested Condition

	var point7 int = 10

	if point7 > 7 {
		switch point7 {
		case 10:
			fmt.Println("Perfect!")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if point7 == 5 {
			fmt.Println("Not bad")
		} else if point == 3{
			fmt.Println("Keep trying")
		} else {
			fmt.Println("You can do it")
			if point7 == 0 {
				fmt.Println("Try harder!")
			}
		}
	}
}