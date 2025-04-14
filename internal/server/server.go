package server

import (
	"Modules/internal/common"
	"Modules/internal/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type serverInterface interface {
	FunctionB()
}

func FunctionA(b serverInterface) {
	b.FunctionB()
}

func GetAllCars(c *gin.Context) {
	cars := common.ReadCarsFromJson()

	c.JSON(http.StatusOK, cars)
}

func AddNewCar(c *gin.Context) {
	cars := common.ReadCarsFromJson()
	newCar := model.Car{}

	c.ShouldBind(&newCar)

	cars = append(cars, newCar)

	common.SaveCarsInJson(cars)

	fmt.Println("New Car Added:")
	newCar.ShowCarInfo()
}

func UpdateCar(c *gin.Context) {
	cars := common.ReadCarsFromJson()
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

	common.SaveCarsInJson(cars)

	fmt.Println("Car updated:")
	updatedCar.ShowCarInfo()
}

func DeleteCar(c *gin.Context) {
	Cars := common.ReadCarsFromJson()

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

	common.SaveCarsInJson(Cars)

	fmt.Println("Car removed")
	Cars[index].ShowCarInfo()

}

func RunServer() {
	router := gin.Default()

	router.GET("/cars", GetAllCars)
	router.POST("/cars/add", AddNewCar)
	router.POST("/cars/update", UpdateCar)
	router.DELETE("/cars/delete", DeleteCar)

	go func() {
		if err := router.Run(":8080"); err != nil {
			fmt.Println("Server get error", err)
		}
	}()
}
