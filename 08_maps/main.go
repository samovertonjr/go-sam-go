package main

import "fmt"

func main() {
	//Define Map
	emails := make(map[string]string)

	//Assign key
	emails["Sam"] = "sam@sam.com"
	emails["Milia"] = "milia@milia.com"

	//short method
	lastName := map[string]string{"Omentri": "omentri@omentri.com", "Overton": "overton@overton.com"}

	fmt.Println(emails)
	fmt.Println(lastName)
}
