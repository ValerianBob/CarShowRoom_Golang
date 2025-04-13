package usecase

import (
	"Modules/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ShowCarsHTTPHandler(w http.ResponseWriter, r *http.Request) {
	Cars := ReadCarsFromJson()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Cars)
}

func AddNewCarHTTPHandler(w http.ResponseWriter, r *http.Request) {

	newCar := model.Car{}

	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("New Car Added:")
	newCar.ShowCarInfo()

	Cars := ReadCarsFromJson()
	Cars = append(Cars, newCar)
	SaveCarsInJson(Cars)
}

func UpdateCarHTTPHandler(w http.ResponseWriter, r *http.Request) {
	Cars := ReadCarsFromJson()

	updatedCar := model.Car{}

	indexStr := r.URL.Query().Get("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(Cars) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&updatedCar)

	fmt.Println("Car updated:")
	updatedCar.ShowCarInfo()

	Cars[index] = updatedCar
	SaveCarsInJson(Cars)
}

func RemoveCarHTTPHandler(w http.ResponseWriter, r *http.Request) {
	Cars := ReadCarsFromJson()

	indexStr := r.URL.Query().Get("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(Cars) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}

	fmt.Println("Car removed")
	Cars[index].ShowCarInfo()

	Cars = append(Cars[:index], Cars[index+1:]...)

	SaveCarsInJson(Cars)
}
