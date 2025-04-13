package main

import (
	"encoding/json"
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

func SaveCarsInJson(c []Car) {
	jsonData, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		fmt.Println("Error with json save", err)
		return
	}

	err = os.WriteFile("CarsData.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func ReadCarsFromJson() []Car {
	AllCars := []Car{}

	carsData, err := os.ReadFile("CarsData.json")
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

func main() {
	AllCars := ReadCarsFromJson()

	//Console inputs :
	userInput := ""
	continueInput := ""
	index := 0

	for {
		fmt.Println("----CAR-SHOWROOM----")
		fmt.Println()
		fmt.Println(" Actions :")
		fmt.Println("1 - Show all cars")
		fmt.Println("2 - Add new car")
		fmt.Println("3 - Change car info")
		fmt.Println("4 - Remove car from showroom")
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
