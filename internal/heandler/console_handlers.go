package handlers

import (
	"Modules/internal/model"
	"Modules/internal/usecase"
	"fmt"
)

func ShowAllCars(AllCars []model.Car, continueInput string) {
	fmt.Println("")
	usecase.ShowAllCar(AllCars)
	fmt.Println()
	fmt.Println("Press enter to continue :")
	fmt.Scanln(&continueInput)
}
