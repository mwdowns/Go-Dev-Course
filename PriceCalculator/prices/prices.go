package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const pricesFile = "prices.txt"
const fileReaderError = "cannot read file, using defaults"

var defaultPrices = []float64{10, 20, 30}

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
	createdAt         time.Time
}

func (job *TaxIncludedPriceJob) loadData() {
	// get prices from file
	data, err := os.Open(pricesFile)
	if err != nil {
		fmt.Println(fileReaderError)
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(data)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(fileReaderError)
		fmt.Println(err)
		data.Close()
		return
	}
	var prices = make([]float64, len(lines))
	for index, line := range lines {
		price, err2 := strconv.ParseFloat(line, 64)
		if err2 != nil {
			fmt.Println(fileReaderError)
			fmt.Println(err2)
			data.Close()
			return
		}
		prices[index] = price
	}
	// everything successful, override default values
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.loadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxPrice := (1 + job.TaxRate) * price
		result[strconv.FormatFloat(price, 'f', 2, 64)] = strconv.FormatFloat(taxPrice, 'f', 2, 64)
	}
	job.TaxIncludedPrices = result
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	// Input prices is set to a default and will be overridden by the loadData function
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: defaultPrices,
	}
}
