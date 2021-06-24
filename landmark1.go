package main

import (
	"fmt"
	"log"
)

func main() {
	location := geo.Landmark{}
	err := location.SetName("The Googleplex")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-1122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.SetName())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
