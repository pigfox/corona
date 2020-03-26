package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	mortality := 0.05
	population := float64(327e6)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Println("US Population:", fmt.Sprintf("%6.0f", population))
	mortalityPercent := fmt.Sprintf("%0.0f", mortality*100)

	fmt.Println("Average Mortality Rate:", mortalityPercent, "%")

	fmt.Fprintf(w, "\n%s\t%s\t", "%Population\nInfection\nProbability", "Number Dead")

	for probability := float64(0.0); probability <= 100; probability += 5 {
		dead := population * probability / 100 * mortality
		fmt.Fprintf(w, "\n%0.2f\t%6.0f\t", probability, dead)
	}
}
