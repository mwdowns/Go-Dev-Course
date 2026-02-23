package cmdmanager

import (
	"fmt"
)

type CMDManager struct {
}

func New() CMDManager {
	return CMDManager{}
}

func (cm CMDManager) ReadLines() ([]string, error) {
	var prices []string
	fmt.Println("Please enter your prices?")
	for {
		var price string
		fmt.Println("Price?")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cm CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}
