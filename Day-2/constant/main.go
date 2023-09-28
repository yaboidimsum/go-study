/* Constant is a variable that cannot be changed after it has been declared.
Such as PI, the value of PI will always be 3.14. */

package main

import "fmt"

func main() {

	const firstName string = "John"
	const lastName = "Doe"

	fmt.Println(firstName)
	fmt.Println(lastName)

	fmt.Print(firstName + " " + lastName + "\n")

	//Multiple constant declaration
	const (
		square = "Kotak"
		isTodayFriday = true
		numeric uint8 = 10
		floatNum = 3.14
	)

	fmt.Println(square)
	// Multiple constant with heritance
	const (
		today string = "Friday"
		thisDay
	)

	fmt.Println(thisDay)

	// Multiple Constant in one line
	const one, two, three = 1, 2, 3

}