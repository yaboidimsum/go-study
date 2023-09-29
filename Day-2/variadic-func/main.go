package main

import (
	"fmt"
	"strings"
)

//Variadic function is a function that can accept multiple arguments. The parameter of a variadic function is a slice. We can pass multiple arguments to a variadic function by separating them with commas.

func main(){
	var numberList = []int{2, 3, 2, 4, 3, 5}
	var age = calculate(numberList...)
	var msg = fmt.Sprintf("Rata-rata umur: %.2f", age)
	fmt.Println(msg)

	var hobbies = []string{"Reading", "Coding", "Swimming"}
	yourHobbies("Rizky", hobbies...)
}

func calculate(numbers... int) float64 {
	var total int = 0

	for _, number := range numbers {
		total += number
	}
	
	var avg = float64(total) / float64(len(numbers))
	fmt.Println("Total\t:", total)
	fmt.Println("Avg\t:", avg)
	return avg
}

func yourHobbies (name string, hobbies ...string){
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}