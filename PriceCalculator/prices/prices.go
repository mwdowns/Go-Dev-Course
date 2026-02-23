package prices

import (
	"fmt"
	"strconv"
	"time"

	"mwdowns.me/price-calculator/converter"
	"mwdowns.me/price-calculator/iomanager"
)

const fileReaderError = "cannot read file, using defaults"

var defaultPrices = []float64{10, 20, 30}

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	IOManager         iomanager.Manager `json:"-"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
	CreatedAt         time.Time         `json:"-"`
}

func (job *TaxIncludedPriceJob) loadData() ([]float64, error) {
	// get prices from file
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return nil, err
	}

	floats, err := converter.StringsToFloats(lines)
	if err != nil {
		return nil, err
	}
	return floats, nil
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
	job.IOManager.WriteResult(job)
	return nil
}

func NewTaxIncludedPriceJob(m iomanager.Manager, taxRate float64) *TaxIncludedPriceJob {
	// Input prices is set to a default and will be overridden by the loadData function
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		IOManager:   m,
		InputPrices: defaultPrices,
		CreatedAt:   time.Now(),
	}
}
