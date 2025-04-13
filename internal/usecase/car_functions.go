package usecase

import (
	"Modules/internal/model"
	"fmt"
)

func ShowAllCar(cars []model.Car) {
	for i := 0; i < len(cars); i++ {
		fmt.Print(i + 1)
		cars[i].ShowCarInfo()
	}
}

func NewCarInfo() model.Car {
	newBrand := ""
	newModel := ""
	newTransmission := ""
	newYear := 0
	newHoursePower := 0
	newPrice := 0

	fmt.Printf("Enter new Brand :")
	fmt.Scanln(&newBrand)
	fmt.Printf("Enter new Model :")
	fmt.Scanln(&newModel)
	fmt.Printf("Enter new Transmission :")
	fmt.Scanln(&newTransmission)
	fmt.Printf("Enter new Year :")
	fmt.Scanln(&newYear)
	fmt.Printf("Enter new Hourse Power :")
	fmt.Scanln(&newHoursePower)
	fmt.Printf("Enter new Price :")
	fmt.Scanln(&newPrice)
	fmt.Printf("")

	newCar := model.Car{
		Brand:        newBrand,
		Model:        newModel,
		Transmission: newTransmission,
		Year:         newYear,
		HoursePower:  newHoursePower,
		Price:        newPrice,
	}

	return newCar
}
