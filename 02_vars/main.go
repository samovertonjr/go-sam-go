package main

import "fmt"

func main() {

	//using var
	name := "Sam"
	var age int32 = 28
	const isCool bool = true
	var size float32 = 2.4

	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", age)
}
