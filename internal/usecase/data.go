package usecase

import (
	"CarShowRoom/internal/model"
	"encoding/json"
	"fmt"
	"os"
)

type JsonDatabase struct {
	jsonData []byte
	err      error
}

func (jd JsonDatabase) SaveCarsInJson(c []model.Car) {
	jd.jsonData, jd.err = json.MarshalIndent(c, "", " ")
	if jd.err != nil {
		fmt.Println("Error with json save", jd.err)
		return
	}

	jd.err = os.WriteFile("internal/adapters/repository/CarsData.json", jd.jsonData, 0644)
	if jd.err != nil {
		fmt.Println("Error writing to file:", jd.err)
		return
	}
}

func (jd JsonDatabase) ReadCarsFromJson() []model.Car {
	AllCars := []model.Car{}

	carsData, err := os.ReadFile("internal/adapters/repository/CarsData.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		AllCars = []model.Car{}
		return AllCars
	}

	err = json.Unmarshal(carsData, &AllCars)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		AllCars = []model.Car{}
		return AllCars
	}

	return AllCars
}
