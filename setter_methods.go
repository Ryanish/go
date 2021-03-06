package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

//Set function with initial letter of type (d) and name of struct with a pointer*, then create METHOD with the param name that will hold a variable from the struct method
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	// Set the struct method type by assigning the accessed value (d.Year) to a variable (year)
	d.Year = year
	return nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	date := Date{}

	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year)

	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Month)

	err = date.SetDay(2)
	if err != nil {
		log.Fatal(err)
	}
	date.Day = 50

	fmt.Println(date)
}
