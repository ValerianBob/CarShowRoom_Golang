package main

import "fmt"

type Car struct {
	Brand        string
	Model        string
	Transmission string
	Year         int
	HoursePower  int
	Price        int
}

func (c Car) ShowCarInfo() {
	fmt.Printf("{Brand : %s, Model : %s, Transmission : %s, Year : %d, Hourse Power : %d HP, Price : %d $}",
		c.Brand, c.Model, c.Transmission, c.Year, c.HoursePower, c.Price)
	fmt.Println("")
}
