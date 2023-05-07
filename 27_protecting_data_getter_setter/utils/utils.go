package utils

import (
	"errors"
	"time"
)

type Date struct {
	// fields are lowercase, they are not accessible outside of the package
	day   int
	month int
	year  int
}

// Setter functions to change the day
// /!\ function names are capitalized, so that they available ouside of the package
func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("incorrect day value")
	}
	d.day = day
	return nil
}

// Setter function to chnage the month
func (d *Date) SetMonth(month int) error {
	if (month < 1) || (month > 12) {
		return errors.New("incorrect month value")
	}
	d.month = month
	return nil
}

// Setter function to chnage the year
func (d *Date) SetYear(year int) error {
	if (year < 1875) || (year > time.Now().Year()) {
		return errors.New("incorrect year value")
	}
	d.year = year
	return nil
}

// Getter functions
func (d *Date) GetDay() int {
	return d.day
}

// Getter functions
func (d *Date) GetMonth() int {
	return d.month
}

// Getter functions
func (d *Date) GetYear() int {
	return d.year
}
