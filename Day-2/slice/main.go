/* Slice is reference element of array. It works like a pointer to an array.*/

package main

import "fmt"

func main() {

	// Slice Declaration doesn't need length
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Print("Isi slice fruits ", fruits)

	fmt.Println("")

	// Slice Initialization
	var fruits1 = []string{"apple", "grape", "banana", "melon"}
	var aFruit = fruits1[0]

	fmt.Println("aFruit is", aFruit)

	fmt.Println("")

	// Array and Slice
	// Array is a bag of values with the same type. Slice is a segment of an array.

	var fruits2 = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits2[0:2]

	fmt.Println("fruits2", newFruits)

	// Slice is a reference to an array. When we change the value of a slice, the value of the array will also change. And vice versa.

	var fruits3 = []string{"apple", "grape", "banana", "melon"}

	var aFruit1 = fruits3[0:3]
	var bFruit1 = fruits3[1:4]

	var aaFruit1 = aFruit1[1:2]
	var baFruit1 = bFruit1[0:1]

	fmt.Println(fruits3)
	fmt.Println(aFruit1)
	fmt.Println(bFruit1)
	fmt.Println(aaFruit1)
	fmt.Println(baFruit1)
	
	fruits3[1] = "papaya"
	fmt.Print("After change index 0 fruits3 ", fruits3)

	fmt.Println(fruits3)
	fmt.Println(aFruit1)
	fmt.Println(bFruit1)
	fmt.Println(aaFruit1)
	fmt.Println(baFruit1)
	fmt.Println(len(fruits3))

	// Cap is the maximum length of a slice. Cap is the length of the underlying array starting from the index of the slice.

	var fruits4 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits4))
	fmt.Println(cap(fruits4))
	fmt.Println("")
	var afruits4 = fruits4[0:3]
	fmt.Println(len(afruits4))
	fmt.Println(cap(afruits4))
	fmt.Println("")
	var bfruits4 = fruits4[1:4]
	fmt.Println(len(bfruits4))
	fmt.Println(cap(bfruits4))

	//Space
	fmt.Println("")

	// Append is used to add elements to a slice. If the length of the slice is not enough, a new array will be created with a larger capacity.

	var fruits5 = []string{"apple", "grape", "banana"}

	var cfruits5 = append(fruits5, "papaya")

	fmt.Println(fruits5)
	fmt.Println(cfruits5)

	//Space
	fmt.Println("")

	var fruits6 = []string{"apple", "grape", "banana"}
	var cfruits6 = fruits6[0:2]

	var dfruits6 = append(cfruits6, "papaya")

	fmt.Println(fruits6)
	fmt.Println(cfruits6)
	fmt.Println(dfruits6)


	//Copy is used to copy the contents of a slice to another slice. The destination slice will be created with a capacity that is enough to accommodate the number of elements copied.

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n:= copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	//Space
	fmt.Println("")

	dst1:= []string{"potato", "potato", "potato"}
	src1:= []string{"watermelon", "pinnaple"}

	n1:= copy(dst1, src1)

	fmt.Println(dst1)
	fmt.Println(src1)
	fmt.Println(n1)

	//Space
	fmt.Println("")

	// Accessing Slice Element with 3 index
	var fruits7 = []string{"apple", "grape", "banana", "melon"}
	var aFruit7 = fruits7[0:2]
	var bFruit7 = fruits7[0:2:2] // The third index is the capacity of the slice

	fmt.Println(fruits7)
	fmt.Println(len(fruits7))
	fmt.Println(cap(fruits7))

	fmt.Println(aFruit7)
	fmt.Println(len(aFruit7))
	fmt.Println(cap(aFruit7))

	fmt.Println(bFruit7)
	fmt.Println(len(bFruit7))
	fmt.Println(cap(bFruit7))
	

}