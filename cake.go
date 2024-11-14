package main

type Cake struct {
	name     string
	category string
	price    float64
}

func (cake *Cake) getPrice() float64 {
	return cake.price
}

func (cake *Cake) setPrice(price float64) {
	cake.price = price
}

func (cake *Cake) setName(name string) {
	cake.name = name
}

func (cake *Cake) getName() string {
	return cake.name
}

func (cake *Cake) setCategory(cakeType string) {
	cake.category = cakeType
}

func (cake *Cake) getCategory() string {
	return cake.category
}
