package model

import "fmt"

type Car struct {
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Transmission string `json:"transmission"`
	Year         int    `json:"year"`
	HoursePower  int    `json:"hoursepower"`
	Price        int    `json:"price"`
}

func (c Car) ShowCarInfo() {
	fmt.Printf("{Brand : %s, Model : %s, Transmission : %s, Year : %d, Hourse Power : %d HP, Price : %d $}",
		c.Brand, c.Model, c.Transmission, c.Year, c.HoursePower, c.Price)
	fmt.Println("")
}
