package main

import (
	"flag"
	"fmt"
)

var income int

func init() {
	const (
		defaultValue = 1000
		usage        = "Income pre-taxes"
	)

	flag.IntVar(&income, "income", defaultValue, usage)
	flag.IntVar(&income, "i", defaultValue, usage+" (shorthand)")
}

type band struct {
	amount  int
	maxTax  int
	rate    int
	whenMax int
}

func main() {
	flag.Parse()

	var bands [4]band

	bands[0] = band{amount: 12500, rate: 0, whenMax: 12500, maxTax: 0}
	bands[1] = band{amount: 50000, rate: 20, whenMax: 30000, maxTax: 7500}
	bands[2] = band{amount: 150000, rate: 40, whenMax: 60000, maxTax: 40000}
	bands[3] = band{amount: 0x7FF0000000000000, rate: 45}

	postTax := 0
	taxPaid := 0
	for i := 0; i < len(bands); i++ {
		band := bands[i]

		// handle the first band as a special case
		if i == 0 {
			if income <= band.amount {
				postTax += income

				break
			} else {
				postTax += band.whenMax

				continue
			}
		}

		if income > band.amount {
			taxPaid += band.maxTax
			postTax += band.whenMax

			continue
		} else {
			prevBand := bands[i-1]

			tax := ((income - prevBand.amount) * band.rate) / 100
			taxPaid += tax
			postTax += (income - prevBand.amount) - tax

			break
		}
	}

	fmt.Printf("Income: %d\n", income)
	fmt.Printf("Tax: %d\n", taxPaid)
	fmt.Printf("Income post-tax: %d\n", postTax)
}
