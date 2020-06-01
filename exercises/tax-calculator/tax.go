package main

import (
	"flag"
	"fmt"
)

var income float64

func init() {
	const (
		defaultValue = 1000
		usage        = "Income pre-taxes"
	)

	flag.Float64Var(&income, "income", defaultValue, usage)
	flag.Float64Var(&income, "i", defaultValue, usage+" (shorthand)")
}

type band struct {
	amount  float64
	maxTax  float64
	rate    float64
	whenMax float64
}

func main() {
	flag.Parse()

	var bands [4]band

	bands[0] = band{amount: 12500, rate: 0, whenMax: 12500, maxTax: 0}
	bands[1] = band{amount: 50000, rate: 20, whenMax: 30000, maxTax: 7500}
	bands[2] = band{amount: 150000, rate: 40, whenMax: 60000, maxTax: 40000}
	bands[3] = band{amount: 0x7FF0000000000000, rate: 45}

	var postTax float64
	var taxPaid float64
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

	fmt.Printf("Income: £%0.2f\n", income)
	fmt.Printf("Tax: £%0.2f\n", taxPaid)
	fmt.Printf("Income post-tax: £%0.2f\n", postTax)
}
