package main

import (
	"fmt"
	"protobuff/protobuff/compiled/testingproto/basic"
)

func main() {

	message := basic.Building{
		BuildingName:   "marvel Studio",
		BuildingNumber: 101,
		Street: &basic.Street{
			Name: "Route 001",
			City: &basic.City{
				Name:        "New York City",
				ZipCode:     "10101",
				CountryName: "United States of America",
			},
		},
	}

	fmt.Println("Hello Proto buffer 3", message)
}
