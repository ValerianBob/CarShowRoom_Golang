package run

import (
	"Modules/internal/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	router.GET("/cars", handler.GetAllCars)
	router.POST("/cars/add", handler.AddNewCar)
	router.PUT("/cars/update", handler.UpdateCar)
	router.DELETE("/cars/delete", handler.DeleteCar)

	go func() {
		if err := router.Run(":8080"); err != nil {
			fmt.Println("Server get error", err)
		}
	}()
}
