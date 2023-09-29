package main

import "fmt"

// Map is a data structure that holds key-value pairs. The key is used to identify the value. The key and value can be of any type as long as the key is unique.


func main() {

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])

	chicken["mei"] = 60

	fmt.Println("mei", chicken["mei"])

	// Map Initialization
	var chicken1 = map[string]int{"januari": 50, "februari": 40}
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	fmt.Println("chicken1", chicken1)
	fmt.Println("chicken2", chicken2)

	// Another way to initialize a map
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)

	fmt.Println("chicken3", chicken3)
	fmt.Println("chicken4", chicken4)
	fmt.Println("chicken5", chicken5)

	// Iterating over a map with for-range

	var chicken6 = map[string]int{"januari": 50, "februari": 40, "maret": 34, "april": 67}

	for key, val:=range chicken6 {
		fmt.Println(key, "  \t:", val)
	}

	// Delete a map element
	var chicken7 = map[string]int{"januari": 50, "februari": 40, "maret": 34, "april": 67}

	fmt.Println(len(chicken7))
	fmt.Println(chicken7)

	delete(chicken7, "januari")

	fmt.Println(len(chicken7))
	fmt.Println(chicken7)

	// Check if a map element exists
	var chicken8 = map[string]int{"januari": 50, "februari": 40, "maret": 34, "april": 67}
	var value, isExist = chicken8["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
	
	// Slice and Map Combinaton

	var chickens = []map[string]string{
		{"name":"chicken blue", "gender":"male"},
		{"name": "chicken red",    "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	// Different key-value types

	var data = []map[string]string{
				{"name": "Rizky", "grade": "A"},
				{"name": "Khairunnisa", "grade": "B"},
				{"name": "Putri", "grade": "A"},
				{"name": "Golang", "community": "A"},
			}

			for _, val:= range data {
				fmt.Println(val["name"], val["grade"], val["community"])
			}
		}