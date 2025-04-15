package console

import (
	"CarShowRoom/internal/model"
	"CarShowRoom/internal/usecase"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/C", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ShowAllCarsConsole(AllCars []model.Car, continueInput string) {
	car_interface := model.Car{}
	fmt.Println("")
	car_interface.ShowAllCar(AllCars)
	fmt.Println()
	fmt.Println("Press enter to continue :")
	fmt.Scanln(&continueInput)
}

func AddCarConsole(AllCars []model.Car, continueInput string) []model.Car {
	data_interface := usecase.JsonDatabase{}

	car_interface := model.Car{}
	fmt.Println("")
	newCar := car_interface.NewCarInfo()
	AllCars = append(AllCars, newCar)

	fmt.Println()
	fmt.Println("Car added !")
	fmt.Println()
	fmt.Println("Press enter to continue :")
	data_interface.SaveCarsInJson(AllCars)
	fmt.Scanln(&continueInput)

	return AllCars
}

func UpdateCarConsole(AllCars []model.Car, continueInput string, index int) []model.Car {
	data_interface := usecase.JsonDatabase{}

	car_interface := model.Car{}
	fmt.Println("")
	car_interface.ShowAllCar(AllCars)
	fmt.Println()
	fmt.Println("Enter index of car to change info :")
	fmt.Scanln(&index)

	if index < 1 || index > len(AllCars) {

		fmt.Println("Wrong index ! Press enter to continue :")
		fmt.Scanln(&continueInput)
		ClearConsole()

	} else {

		newCar := car_interface.NewCarInfo()
		AllCars[index-1] = newCar

		fmt.Println()
		fmt.Println("Car Info Updated !")
		fmt.Println()
		fmt.Println("Press enter to continue :")
		data_interface.SaveCarsInJson(AllCars)
		fmt.Scanln(&continueInput)
		ClearConsole()
	}
	return AllCars
}

func DeleteCarConsole(AllCars []model.Car, continueInput string, index int) []model.Car {
	data_interface := usecase.JsonDatabase{}

	car_interface := model.Car{}
	if len(AllCars) == 1 {

		fmt.Println()
		fmt.Println("You can't remove last car. Press enter to continue :")
		fmt.Scanln(&continueInput)
		ClearConsole()

	} else {

		fmt.Println("")

		car_interface.ShowAllCar(AllCars)

		fmt.Println()
		fmt.Println("Enter index of car to remove :")
		fmt.Scanln(&index)

		if index < 1 || index > len(AllCars) {

			fmt.Println("Wrong index ! Press enter to continue :")
			fmt.Scanln(&continueInput)
			ClearConsole()

		} else {

			AllCars = append(AllCars[:index-1], AllCars[index:]...)

			fmt.Println()
			fmt.Println("Car removed !")
			fmt.Println()
			fmt.Println("Press enter to continue :")
			data_interface.SaveCarsInJson(AllCars)
			fmt.Scanln(&continueInput)
			ClearConsole()
		}
	}
	return AllCars
}
