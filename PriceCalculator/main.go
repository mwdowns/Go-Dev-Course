package main

import (
	"fmt"

	"mwdowns.me/price-calculator/cmdmanager"
	"mwdowns.me/price-calculator/filemanager"
	"mwdowns.me/price-calculator/prices"
)

const pricesFileName = "prices.txt"

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		cmdm := cmdmanager.New()
		fm := filemanager.New(pricesFileName, fmt.Sprintf("job_%.0f.json", taxRate*100))
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		job2 := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := job.Process()
		err2 := job2.Process()
		if err != nil || err2 != nil {
			return
		}
	}
	fmt.Println("Finished")
}
