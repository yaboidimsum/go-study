package main

import "fmt"

func main (){
	var firstName string = "John"
	var lastName string

	lastName = "Doe"

	fmt.Printf("Hello %s %s!\n", firstName, lastName)

	// Deklarasi Variabel tanpa tipe data
	var fullName = "John Doe"
	fmt.Println("Hello", fullName + "!")

	// Deklarasi Variabel tanpa kata kunci var
	age:=27
	fmt.Println("My age is", age)
	//Jangan pakai := lagi, karena sudah dideklarasikan sebelumnya
	age = 29
	fmt.Println("My age is", age)

	//Multiple Variabel Declaration
	var name1, name2, name3 string = "John", "Joe", "Smith"
	fmt.Println("Name 1:", name1, "Name 2:", name2, "Name 3:", name3)

	//underscore variable
	_ = "Belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "John", "Doe"

	fmt.Println("Name:", name)

	// Deklarasi Variable menggunakan keyword new
	nameNew := new(string)

	fmt.Println(nameNew)
	fmt.Println(*nameNew)

}