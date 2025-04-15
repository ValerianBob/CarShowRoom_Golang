package run

import (
	"Modules/internal/adapters/handler/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	cars := router.Group("/cars")
	{
		cars.GET("/", http.GetAllCars)
		cars.GET("/:id", http.GetCarById)
		cars.POST("/", http.AddNewCar)
		cars.PUT("/:id", http.UpdateCar)
		cars.DELETE("/:id", http.DeleteCar)
	}

	go func() {
		if err := router.Run(":8080"); err != nil {
			fmt.Println("Server get error", err)
		}
	}()
}
