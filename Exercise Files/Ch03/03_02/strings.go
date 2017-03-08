package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicitly typed string"
	fmt.Printf("str1: %v:%T\n", str1, str1)
	fmt.Println(str1)

	str2 := "An explicitly typed string"
	fmt.Printf("str2: %v:%T\n", str2, str2)

	// to uppercase
	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	// compare strings
	lvalue := "hello"
	uvalue := "HELLO"
	fmt.Println("Equal?", (lvalue == uvalue))
	fmt.Println("Equal non-case-sensitive?", strings.EqualFold(lvalue, uvalue))

	// does one string contain another?
	fmt.Println("Contains exp?", strings.Contains(str1, "exp"))
	fmt.Println("Contains exp?", strings.Contains(str2, "exp"))
}
