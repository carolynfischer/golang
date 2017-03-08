package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(strings.ToUpper("Hello, World!"))

	str1 := "The quick brown fox"
	str2 := "jumped over"
	str3 := "the lazy dog."
	aNumber := 42
	isTrue := true

	stringlength, err := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("String length:", stringlength)
	}

	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of aNumber: %v\n", isTrue)
	fmt.Printf("Value of aNumber: %.2f\n", float64(aNumber))

	// print type
	fmt.Printf("Data types: %T, %T, %T, %T and %T\n", str1, str2, str3, aNumber, isTrue)
	myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T and %T\n", str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)
}
