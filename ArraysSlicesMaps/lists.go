package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	//array - fixed number of entries
	prices := [4]float64{10.99, 20.0, 14.99, 5.99}
	fmt.Println(prices)
	//slice - dynamic number of entries
	var productNames []string
	fmt.Println(productNames)
	//creating a slice from an array
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	featuredPrices = append(featuredPrices, 49.99)
	fmt.Println(featuredPrices)

	product1 := Product{id: uuid.New().String(), title: "Product 1", price: 100}
	product2 := Product{id: uuid.New().String(), title: "Product 2", price: 200}

	products := []Product{product1, product2}

	for _, product := range products {
		fmt.Println(product.id, product.title)
	}
}
