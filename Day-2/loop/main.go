/* Go only provides one looping construct, the for loop. But it combines for, foreach, and while of other languages.*/

package main

import "fmt"

func main() {
// For loop
for i := 0; i < 5; i++ {
	fmt.Println("Angka", i)
	}
	fmt.Print("\n")

// For loop like while

	var j int  = 0

	for j < 5 {
		fmt.Print("Angka", j, "\n")
		j++
	}

	fmt.Print("\n")
	
	// For without argument
	var k int = 0
	
	for {
		fmt.Println("Angka", k)
		k++
		if k == 5 {
			break
		}
	}
	fmt.Print("\n")


	// For-Range

	var xs = "123"

	for i, v := range xs {
		fmt.Printf("index ke %d = %d\n", i, v)
	}
	fmt.Print("\n")

	var ys = [5]int{1, 2, 3, 4, 5} //Array
	for _, v := range ys {
		fmt.Printf("value = %d\n", v)
	}
	fmt.Print("\n")

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
    fmt.Println("Key=", k, "Value=", v)
	}

	// Break and Continue
	for l := 1; l <= 10; l++ {
		if l % 2 == 1 {
			continue
		}
		if l > 8 {
			break
		}
		fmt.Println("Angka", l)
	}
	fmt.Print("\n")

	// Nested Loop
	for m := 0; m < 5; m++ {
		for n := m; n < 5; n++ {
			fmt.Print(n, " ")
		}
		fmt.Println()
	}
	//Outer loop
	outerLoop:
	for o := 0; o < 5; o++ {
		for p := o; p < 5; p++ {
			if p == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", o, "][", p, "]", "\n")
		}
	}
}
