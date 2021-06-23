package main

import (
	"fmt"
	"log"

	"github.com/ryanish/go/coordinates"
)

func main() {
	coordinates :=
	err := 	coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Lat)
}