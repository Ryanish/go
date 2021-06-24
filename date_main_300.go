package main1

import (
	"fmt"
	"log"

	"github.com/ryanish/go/calendar"
)

func main1() {

	event := calendar.Event{}
	err := event.SetTitle("Mum's birthday!")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	// promoted from the original Date type in the Event struct
	err = event.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())

	// date := calendar.Date{}
	// err := date.SetYear(2019)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = date.SetMonth(12)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = date.SetDay(27)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(date.Year())
	// fmt.Println(date.Month())
	// fmt.Println(date.Day())
}
