package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	//Assumed average mortality rate
	mortality := 0.013

	//US population
	population := float64(327e6)

	//African/Middle East/Asia
	//population := float64(1e9)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Println("Population:", fmt.Sprintf("%6.0f", population))
	mortalityPercent := fmt.Sprintf("%0.2f", mortality*100)

	fmt.Println("Average Mortality Rate:", mortalityPercent, "%")

	fmt.Fprintf(w, "\n%s\t%s\t%s\t", "%Population\nInfection\nProbability", " | ", "Number Dead")

	//Iterate over probabilities to calculate number of dead
	for probability := float64(0.0); probability <= 100; probability += 5 {
		dead := population * probability / 100 * mortality
		fmt.Fprintf(w, "\n%2.f\t%s\t%9.f\t", probability, " | ", dead)
	}
}
