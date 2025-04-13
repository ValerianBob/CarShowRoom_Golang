package main

import "fmt"

type Car struct {
	Brand        string
	Model        string
	Transmission string
	Year         int
	HoursePower  int
	Price        int
}

func (c Car) ShowCarInfo() {
	fmt.Printf("{Brand : %s, Model : %s, Transmission : %s, Year : %d, Hourse Power : %d HP, Price : %d $}",
		c.Brand, c.Model, c.Transmission, c.Year, c.HoursePower, c.Price)
	fmt.Println("")
}

func ShowAllCar(cars []Car) {
	for i := 0; i < len(cars); i++ {
		fmt.Print(i + 1)
		cars[i].ShowCarInfo()
	}
}

func NewCarInfo() Car {
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

	newCar := Car{
		Brand:        newBrand,
		Model:        newModel,
		Transmission: newTransmission,
		Year:         newYear,
		HoursePower:  newHoursePower,
		Price:        newPrice,
	}

	return newCar
}
