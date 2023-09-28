/* Array is a collection of data with the same type, which is fixed in size and has a fixed address in memory.*/

package main

import "fmt"


func main() {
	var names [4]string
	names[0] = "Rizky"
	names[1] = "Khairunnisa"
	names[2] = "Putri"
	names[3] = "Salsabila"
	
	fmt.Println(names[0], names[1], names[2], names[3])

	// Array Initialization
	var fruits = [5]string {"apple", "grape", "banana", "melon", "orange"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// Array initialization without length

	var numbers = [...]int{2, 3, 2, 4, 3, 5}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// Array multidimensional
	var numbers1 = [3][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}, [3]int{5, 6, 7}}
	var numbers2 = [3][3]int{{3, 2, 3}, {3, 4, 5}, {5, 6, 7}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// Array Element with For
	var fruits1 = [4]string{"apple", "grape", "banana", "melon"}

	for i:=0; i<len(fruits1); i++ {
		fmt.Printf("element %d : %s\n", i, fruits1[i])
	}
	fmt.Println("")
	// Array Element with For-Range
	var fruits2 = [4]string{"apple", "grape", "banana", "melon"}

	for j, fruit:= range fruits2 {
		fmt.Printf("element %d : %s\n", j, fruit)
	}

	fmt.Println("")

	// _ is used to ignore index
	var fruits3 = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit:= range fruits3 {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	fmt.Println("")

	// Array allocation with make
	var fruits4 = make([]string, 2)

	fruits4[0] = "apple"
	fruits4[1] = "manggo"

	fmt.Println(fruits4)
}