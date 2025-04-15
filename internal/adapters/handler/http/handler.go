package http

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

func GetCarById(c *gin.Context) {
	cars := usecase.ReadCarsFromJson()

	indexStr := c.Param("id")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Println("Not correct id")
		return
	}

	carById := cars[index]

	c.JSON(http.StatusOK, carById)
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

	indexStr := c.Param("id")
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
	cars := usecase.ReadCarsFromJson()

	indexStr := c.Param("id")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Println("Index is not correct !")
		return
	}

	if index < 0 || index >= len(cars) {
		fmt.Println("Index not correct !")
		return
	}

	cars = append(cars[:index], cars[index+1:]...)

	usecase.SaveCarsInJson(cars)

	fmt.Println("Car removed")
	cars[index].ShowCarInfo()
}
