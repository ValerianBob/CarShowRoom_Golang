package usecase

import (
	"Modules/internal/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func ShowCarsHTTPHandler(w http.ResponseWriter, r *http.Request) {
// 	Cars := ReadCarsFromJson()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(Cars)
// }

func GetAllCars(c *gin.Context) {
	cars := ReadCarsFromJson()

	c.JSON(http.StatusOK, cars)
}

// func AddNewCarHTTPHandler(w http.ResponseWriter, r *http.Request) {

// 	newCar := model.Car{}

// 	err := json.NewDecoder(r.Body).Decode(&newCar)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println("New Car Added:")
// 	newCar.ShowCarInfo()

// 	Cars := ReadCarsFromJson()
// 	Cars = append(Cars, newCar)
// 	SaveCarsInJson(Cars)
// }

func AddNewCar(c *gin.Context) {
	cars := ReadCarsFromJson()
	newCar := model.Car{}

	c.ShouldBind(&newCar)

	cars = append(cars, newCar)

	SaveCarsInJson(cars)

	fmt.Println("New Car Added:")
	newCar.ShowCarInfo()
}

// func UpdateCarHTTPHandler(w http.ResponseWriter, r *http.Request) {
// 	Cars := ReadCarsFromJson()

// 	updatedCar := model.Car{}

// 	indexStr := r.URL.Query().Get("index")
// 	index, err := strconv.Atoi(indexStr)
// 	if err != nil || index < 0 || index >= len(Cars) {
// 		http.Error(w, "Invalid index", http.StatusBadRequest)
// 		return
// 	}

// 	json.NewDecoder(r.Body).Decode(&updatedCar)

// 	fmt.Println("Car updated:")
// 	updatedCar.ShowCarInfo()

// 	Cars[index] = updatedCar
// 	SaveCarsInJson(Cars)
// }

func UpdateCar(c *gin.Context) {
	cars := ReadCarsFromJson()
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

	SaveCarsInJson(cars)

	fmt.Println("Car updated:")
	updatedCar.ShowCarInfo()
}

// func RemoveCarHTTPHandler(w http.ResponseWriter, r *http.Request) {
// 	Cars := ReadCarsFromJson()

// 	indexStr := r.URL.Query().Get("index")
// 	index, err := strconv.Atoi(indexStr)
// 	if err != nil || index < 0 || index >= len(Cars) {
// 		http.Error(w, "Invalid index", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println("Car removed")
// 	Cars[index].ShowCarInfo()

// 	Cars = append(Cars[:index], Cars[index+1:]...)

// 	SaveCarsInJson(Cars)
// }

func DeleteCar(c *gin.Context) {
	Cars := ReadCarsFromJson()

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

	SaveCarsInJson(Cars)

	fmt.Println("Car removed")
	Cars[index].ShowCarInfo()

}

func RunServer() {
	router := gin.Default()

	router.GET("/cars", GetAllCars)
	router.POST("/cars/add", AddNewCar)
	router.POST("/cars/update", UpdateCar)
	router.DELETE("/cars/delete", DeleteCar)

	// http.HandleFunc("/cars", ShowCarsHTTPHandler)
	// http.HandleFunc("/cars/add", AddNewCarHTTPHandler)
	// http.HandleFunc("/cars/update", UpdateCarHTTPHandler)
	// http.HandleFunc("/cars/delete", RemoveCarHTTPHandler)

	// router.Run(":8080")
	go func() {
		if err := router.Run(":8080"); err != nil {
			fmt.Println("Server get error", err)
		}
	}()
}
