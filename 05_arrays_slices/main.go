package main

import (
	"fmt"
)

func main() {
	//Arrays

	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	//or

	peopleArr := [2]string{"sam", "milia"}

	//Slices

	veggieSlice := [3]string{"carrot", "spinach", "brocolli"}

	fmt.Println(fruitArr)
	fmt.Println(peopleArr)
	fmt.Println(veggieSlice)
}
