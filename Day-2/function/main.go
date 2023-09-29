package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	var names = []string{"Khairunnisa", "Putri", "Salsabila"}
	printMessage("Hallo", names)

	var randomValue int
	
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	var randomizer = rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println("random number:", randomizer.Intn(100))

	fmt.Println(divideNum(10,2))
	divideNum(4,0)

}

func printMessage(message string, name []string) {
	var nameString = strings.Join(name, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = rand.Int() % (max - min + 1) + min
	return value
}

// Parameter declaration with same type

func nameOfFuncStr(paramA, paramB string) {
	// ...
}
func nameOfFuncInt(paramA, paramB, paramC int) {
	// ...
}

// Return to store the result of a function

func divideNum(m, n int) int {
	if n == 0{
		fmt.Println("Divider can't be zero")
		return 0	
	}
	var res = m / n
	return res
}