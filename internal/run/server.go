package run

import (
	"Modules/internal/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	router.GET("/cars", handlers.GetAllCars)
	router.POST("/cars/add", handlers.AddNewCar)
	router.PUT("/cars/update", handlers.UpdateCar)
	router.DELETE("/cars/delete", handlers.DeleteCar)

	go func() {
		if err := router.Run(":8080"); err != nil {
			fmt.Println("Server get error", err)
		}
	}()
}
