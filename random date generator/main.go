package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	year, leap := year()
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		if leap {
			daysInMonth = 29
		}
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}

func year() (int, bool) {
	leap := false
	t := time.Now()
	year := rand.Intn(t.Year() + 1)
	switch leap {
	case year%400 == 0 || (year%4 == 0 && year%100 != 0):
		leap = true
	}

	return year, leap
}
