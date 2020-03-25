package main

import "fmt"

func main() {
	mortality := 0.05
	population := float64(327e6)

	for probability := float64(0.0); probability <= 100; probability += 5 {
		percent := fmt.Sprintf("%0.2f", probability)
		fmt.Println("With probability of:", percent, "%")

		dead := fmt.Sprintf("%6.2f", population*probability/100*mortality)
		fmt.Println("Number dead:", dead)
	}
}
