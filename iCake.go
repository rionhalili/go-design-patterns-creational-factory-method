package main

type ICake interface {
	setName(name string)
	setCategory(cakeType string)
	setPrice(price float64)
	getName() string
	getCategory() string
	getPrice() float64
}
