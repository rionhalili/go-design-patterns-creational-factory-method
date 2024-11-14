package main

type orange struct {
	Cake
}

func newOrangeCake() ICake {
	return &orange{
		Cake{
			name:     "Orange Cake",
			category: "Fruits",
			price:    8.7,
		},
	}
}
