package usecase

import (
	"fmt"
	"net/http"
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

func Init() {
	AllCars := ReadCarsFromJson()

	//Console inputs :
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
		fmt.Println("HTTP :")
		fmt.Println("5 - Show all cars by HTTP")
		fmt.Println("6 - Add new car by HTTP")
		fmt.Println("7 - Update car by HTTP")
		fmt.Println("8 - Remove car by HTTP")
		fmt.Println()
		fmt.Println("0 - Exit")
		fmt.Println()
		fmt.Println("Enter your action :")
		fmt.Scanln(&userInput)

		switch userInput {

		case "1":
			fmt.Println("")
			ShowAllCar(AllCars)
			fmt.Println()
			fmt.Println("Press enter to continue :")
			fmt.Scanln(&continueInput)
			ClearConsole()

		case "2":
			fmt.Println("")
			newCar := NewCarInfo()
			AllCars = append(AllCars, newCar)

			fmt.Println()
			fmt.Println("Car added !")
			fmt.Println()
			fmt.Println("Press enter to continue :")
			fmt.Scanln(&continueInput)
			SaveCarsInJson(AllCars)
			ClearConsole()

		case "3":
			fmt.Println("")
			ShowAllCar(AllCars)
			fmt.Println()
			fmt.Println("Enter index of car to change info :")
			fmt.Scanln(&index)

			if index < 1 || index > len(AllCars) {

				fmt.Println("Wrong index ! Press enter to continue :")
				fmt.Scanln(&continueInput)
				ClearConsole()

			} else {

				newCar := NewCarInfo()
				AllCars[index-1] = newCar

				fmt.Println()
				fmt.Println("Car Info Updated !")
				fmt.Println()
				fmt.Println("Press enter to continue :")
				fmt.Scanln(&continueInput)
				SaveCarsInJson(AllCars)
				ClearConsole()
			}

		case "4":
			if len(AllCars) == 1 {

				fmt.Println()
				fmt.Println("You can't remove last car. Press enter to continue :")
				fmt.Scanln(&continueInput)
				ClearConsole()

			} else {

				fmt.Println("")

				ShowAllCar(AllCars)

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
					fmt.Scanln(&continueInput)
					SaveCarsInJson(AllCars)
					ClearConsole()
				}
			}

		case "5":
			http.HandleFunc("/cars", ShowCarsHTTPHandler)
			fmt.Println("Server running on http://localhost:8080")
			fmt.Println("Cars data sent on Browser !")
			http.ListenAndServe(":8080", nil)

			fmt.Scanln(&continueInput)
			ClearConsole()

		case "6":
			http.HandleFunc("/add-new-car", AddNewCarHTTPHandler)
			fmt.Println("Server running on http://localhost:8080")
			http.ListenAndServe(":8080", nil)

		case "7":
			http.HandleFunc("/update-car", UpdateCarHTTPHandler)
			fmt.Println("Server running on http://localhost:8080")
			http.ListenAndServe(":8080", nil)

		case "8":
			http.HandleFunc("/remove-car", RemoveCarHTTPHandler)
			fmt.Println("Server running on http://localhost:8080")
			http.ListenAndServe(":8080", nil)

		case "0":
			fmt.Println()
			fmt.Println("Good bye !")
			fmt.Println()
			return

		default:
			fmt.Println()
			fmt.Println("Invalid action. Press enter to continue :")
			fmt.Scanln(&continueInput)
			ClearConsole()
			fmt.Println()
		}
	}
}
