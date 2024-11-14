package main

type apple struct {
	Cake
}

func newAppleCake() ICake {
	return &apple{
		Cake: Cake{
			name:     "Apple",
			category: "Fruits",
			price:    5.5,
		},
	}
}
