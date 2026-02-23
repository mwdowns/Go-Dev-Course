package main

import (
	"fmt"

	"mwdowns.me/price-calculator/filemanager"
	"mwdowns.me/price-calculator/prices"
)

const pricesFileName = "prices.txt"

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New(pricesFileName, fmt.Sprintf("job_%.0f.json", taxRate*100))
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := job.Process()
		if err != nil {
			return
		}
	}
	fmt.Println("Finished")
}
