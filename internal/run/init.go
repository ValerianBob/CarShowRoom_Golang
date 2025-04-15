package run

import (
	"CarShowRoom/internal/adapters/handler/console"
	"CarShowRoom/internal/usecase"
	"fmt"
)

func Init() {

	RunServer()

	data_interface := usecase.JsonDatabase{}

	AllCars := data_interface.ReadCarsFromJson()

	// Console inputs :
	userInput := ""
	continueInput := ""
	index := 0

	for {
		fmt.Println("----CAR-SHOWROOM----")
		fmt.Println()
		fmt.Println(" Actions :")
		fmt.Println("Console :")
		fmt.Println("1 - Show all cars")
		fmt.Println("2 - Add new car")
		fmt.Println("3 - Change car info")
		fmt.Println("4 - Remove car from showroom")
		fmt.Println("You could do same by HTTP :")
		fmt.Println("Show all cars by http://localhost:8080/cars")
		fmt.Println("Add car by http://localhost:8080/cars/add")
		fmt.Println("Update car by http://localhost:8080/cars/update")
		fmt.Println("Delete car by http://localhost:8080/cars/delete")
		fmt.Println()
		fmt.Println("0 - Exit")
		fmt.Println()
		fmt.Println("Enter your action :")
		fmt.Scanln(&userInput)

		switch userInput {

		case "1":
			console.ShowAllCarsConsole(AllCars, continueInput)
		case "2":
			AllCars = console.AddCarConsole(AllCars, continueInput)
		case "3":
			AllCars = console.UpdateCarConsole(AllCars, continueInput, index)
		case "4":
			AllCars = console.DeleteCarConsole(AllCars, continueInput, index)

		case "0":
			fmt.Println()
			fmt.Println("Good bye !")
			fmt.Println()
			return

		default:
			fmt.Println()
			fmt.Println("Invalid action. Press enter to continue :")
			fmt.Scanln(&continueInput)
			console.ClearConsole()
			fmt.Println()
		}
	}
}
