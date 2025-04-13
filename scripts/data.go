package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveCarsInJson(c []Car) {
	jsonData, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		fmt.Println("Error with json save", err)
		return
	}

	err = os.WriteFile("../assets/CarsData.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func ReadCarsFromJson() []Car {
	AllCars := []Car{}

	carsData, err := os.ReadFile("../assets/CarsData.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		AllCars = []Car{}
		return AllCars
	}

	err = json.Unmarshal(carsData, &AllCars)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		AllCars = []Car{}
		return AllCars
	}

	return AllCars
}
