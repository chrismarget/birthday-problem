package main

import (
	"flag"
	"fmt"
	"math"
)

type config struct {
	peoplePerRoom int
	daysPerYear   int
}

func getConfig() (config) {
	peoplePerRoom := flag.Int("p", 10, "How many people in the room? (default 10)")
	daysPerYear := flag.Int("d", 365, "How many days per year? (default 365)")
	flag.Parse()
	return config {
		peoplePerRoom: *peoplePerRoom,
		daysPerYear: *daysPerYear,
	}
}

func main() {
	cfg := getConfig()
	p := float64(1)
	for i := 1; i < cfg.peoplePerRoom;  i++ {
		p = p * float64(cfg.daysPerYear - i) / float64(cfg.daysPerYear)
	}
	p = float64(1-p)
	oneInX := int(math.Round(1/p))
	fmt.Printf("The probability of finding birthday twins (with %d days/year) among %d people is %f or 1 out of %d\n",
		cfg.daysPerYear,
		cfg.peoplePerRoom,
		p,
		oneInX)
}
