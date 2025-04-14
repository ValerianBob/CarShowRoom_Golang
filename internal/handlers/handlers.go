package handlers

import (
	"Modules/internal/model"
	"Modules/internal/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCars(c *gin.Context) {
	cars := usecase.ReadCarsFromJson()

	c.JSON(http.StatusOK, cars)
}

func AddNewCar(c *gin.Context) {
	cars := usecase.ReadCarsFromJson()
	newCar := model.Car{}

	c.ShouldBind(&newCar)

	cars = append(cars, newCar)

	usecase.SaveCarsInJson(cars)

	fmt.Println("New Car Added:")
	newCar.ShowCarInfo()
}

func UpdateCar(c *gin.Context) {
	cars := usecase.ReadCarsFromJson()
	updatedCar := model.Car{}

	indexStr := c.Query("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Println("Index not correct !")
		return
	}

	if index < 0 || index >= len(cars) {
		fmt.Println("Index not correct !")
		return
	}

	c.ShouldBindJSON(&updatedCar)

	cars[index] = updatedCar

	usecase.SaveCarsInJson(cars)

	fmt.Println("Car updated:")
	updatedCar.ShowCarInfo()
}

func DeleteCar(c *gin.Context) {
	Cars := usecase.ReadCarsFromJson()

	indexStr := c.Query("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Println("Index is not correct !")
		return
	}

	if index < 0 || index >= len(Cars) {
		fmt.Println("Index not correct !")
		return
	}

	Cars = append(Cars[:index], Cars[index+1:]...)

	usecase.SaveCarsInJson(Cars)

	fmt.Println("Car removed")
	Cars[index].ShowCarInfo()
}
