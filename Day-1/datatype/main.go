package main

import "fmt"

func main () {

	/*
	Data Type Numerikal Non-Desimal

	uint: Tipe data bilangan bulat positif
	int: Tipe data bilangan bulat positif dan negatif

	uint8	0 ↔ 255
	uint16	0 ↔ 65535
	uint32	0 ↔ 4294967295
	uint64	0 ↔ 18446744073709551615
	uint	sama dengan uint32 atau uint64 (tergantung nilai)
	byte	sama dengan uint8
	int8	-128 ↔ 127
	int16	-32768 ↔ 32767
	int32	-2147483648 ↔ 2147483647
	int64	-9223372036854775808 ↔ 9223372036854775807
	int	sama dengan int32 atau int64 (tergantung nilai)
	rune	sama dengan int32
	
	*/

	var positiveNumber uint8 = 69
	var negativeNumber = -1243423644

	fmt.Printf("Bilangan positif: %d\n", positiveNumber)
	fmt.Printf("Bilangan negatif: %d\n", negativeNumber)

	/*
	Data Type Numerikal Desimal
	float32	-3.4E+38 ↔ 3.4E+38
	float64	-1.8E+308 ↔ 1.8E+308
	float	sama dengan float32 atau float64 (tergantung nilai)
	*/

	var decimalNumber float32 = 2.62

	fmt.Printf("Bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("Bilangan desimal: %.3f\n", decimalNumber)

	/*
	Data Type Boolean
	bool	true atau false
	*/

	var exist bool = true
	fmt.Printf("Exist? %t \n", exist)

	/*
	Nil Type
	nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya nil berarti memiliki nilai kosong.
	*/
}