package prices

import (
	"fmt"
	"strconv"
	"time"

	"mwdowns.me/price-calculator/converter"
	"mwdowns.me/price-calculator/filemanager"
)

const fileReaderError = "cannot read file, using defaults"

var defaultPrices = []float64{10, 20, 30}

type TaxIncludedPriceJob struct {
	TaxRate           float64
	IOManager         filemanager.FileManager
	InputPrices       []float64
	TaxIncludedPrices map[string]string
	CreatedAt         time.Time
}

func (job *TaxIncludedPriceJob) loadData() ([]float64, error) {
	// get prices from file
	data, lines, err := job.IOManager.ReadFile()
	if err != nil {
		return nil, err
	}

	return converter.StringsToFloats(data, lines)
}

func (job *TaxIncludedPriceJob) Process() error {
	inputPrices, err := job.loadData()
	if err != nil {
		fmt.Println(fileReaderError)
		return err
	}
	job.InputPrices = inputPrices
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxPrice := (1 + job.TaxRate) * price
		result[strconv.FormatFloat(price, 'f', 2, 64)] = strconv.FormatFloat(taxPrice, 'f', 2, 64)
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteJson(job)
	return nil
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	// Input prices is set to a default and will be overridden by the loadData function
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		IOManager:   fm,
		InputPrices: defaultPrices,
		CreatedAt:   time.Now(),
	}
}
