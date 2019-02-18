package main

import "fmt"

func main() {
	ids := []int{44, 33, 43, 432, 244}

	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}

	//removed i and replace with _ so I dont have to use the variable
	for _, id := range ids {
		fmt.Printf("ID %d\n", id)
	}
}
