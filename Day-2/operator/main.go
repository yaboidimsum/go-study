package main

import "fmt"

func main() {

	//Arithmetic Operator
	/*
	+ 	= Addition
	- 	= Subtraction
	star	= Multiplication
	/ 	= Division
	% 	= Modulus
	++ 	= Increment
	-- 	= Decrement
	*/
	//Example
	var a int = (((2 + 6) % 3) * 4 - 2) / 3
	fmt.Println(a)

	//Comparison Operator
	/*
	== 	= Equal
	!= 	= Not Equal
	> 	= Greater Than
	< 	= Less Than
	>= 	= Greater Than Equal
	<= 	= Less Than Equal
	*/

	var isEqual bool = (a == 2)
	fmt.Println(isEqual)

	//Logical Operator
	/*
	&& 	= AND
	|| 	= OR
	! 	= NOT
	*/

	var left bool = false
	var right bool = true

	var leftAndRight bool = left && right

	fmt.Println(leftAndRight)

	var leftOrRight bool = left || right

	fmt.Println(leftOrRight)

	var notLeft bool = !left

	fmt.Println(notLeft)
	

}