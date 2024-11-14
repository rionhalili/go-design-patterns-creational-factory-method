package main

type Cake struct {
	name     string
	cakeType string
}

func (cake *Cake) setName(name string) {
	cake.name = name
}

func (cake *Cake) getName() string {
	return cake.name
}

func (cake *Cake) setType(cakeType string) {
	cake.cakeType = cakeType
}

func (cake *Cake) getType() string {
	return cake.cakeType
}
