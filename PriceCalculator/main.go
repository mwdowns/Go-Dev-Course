package main

import (
	"fmt"

	"mwdowns.me/price-calculator/filemanager"
	"mwdowns.me/price-calculator/prices"
)

const pricesFileName = "prices.txt"

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errChans[index] = make(chan error)
		//cmdm := cmdmanager.New()
		fm := filemanager.New(pricesFileName, fmt.Sprintf("job_%.0f.json", taxRate*100))
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		//job2 := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		go job.Process(doneChans[index], errChans[index])
		//err2 := job2.Process(doneChans[index])
		//if err != nil || err2 != nil {
		//if err != nil {
		//	fmt.Println("could not process job")
		//	return
		//}
	}
	for index := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
		}
	}
	fmt.Println("Finished")
}
