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

type productsMap map[string]Product

func (p productsMap) showTitles() {
	for _, product := range p {
		fmt.Println(product.title)
	}
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
	product3 := Product{id: uuid.New().String(), title: "Product 3", price: 300}

	products := []Product{product1, product2}
	newProducts := []Product{product3}
	products = append(products, newProducts...)

	for _, product := range products {
		fmt.Println(product.id, product.title)
	}

	websites := map[string]string{"google": "www.google.com", "amazon": "www.amazon.com"}
	websites["linkedin"] = "www.linkedin.com"
	fmt.Println(websites)

	productStuff := make(productsMap)
	for index, product := range products {
		fmt.Println(index, product.id)
		productStuff[product.id] = product
	}
	fmt.Println(productStuff)
	productStuff.showTitles()

	userNames := make([]string, 2)
	userNames[0] = "Matt"
	userNames[1] = "Jeri"
	fmt.Println(userNames)
}
