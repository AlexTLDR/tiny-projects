/*
Random ticket generator. Generates a random ticket price and duration for a space flight.
*/

package main

import (
	"fmt"
	"math/rand"
)

const secondsPerDay = 86400

func main() {
	distance := 62100000
	company := ""
	trip := ""

	fmt.Println("Spaceline        Days Trip type  Price")
	fmt.Println("======================================")
	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}
		speed := rand.Intn(15) + 16
		duration := distance / speed / secondsPerDay
		price := 20.0 + speed
		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			price = price * 2
		} else {
			trip = "One-way"
		}
		fmt.Printf("%-16v %4v %-10v $%4.0v\n", company, duration, trip, price)
	}
}
