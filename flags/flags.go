package main

import (
	"flag"
	"fmt"
)

var income int

func init() {
	const (
		defaultIncome = 0
		usage         = "income pre-taxes"
	)
	flag.IntVar(&income, "income", defaultIncome, usage)
	flag.IntVar(&income, "i", defaultIncome, usage+" (shorthand)")
}

func main() {
	flag.Parse()

	fmt.Printf("%d\n", income)
}
