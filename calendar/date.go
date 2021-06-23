package calendar

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

//Setter Methods
//Set function with initial letter of type (d) and name of struct with a pointer*, then create METHOD with the param name that will hold a variable from the struct method
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	// Set the struct method type by assigning the accessed value (d.Year) to a variable (year)
	d.year = year
	return nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

//Getter Methods
func (d *Date) Year() int {
	return d.year
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Day() int {
	return d.day
}
