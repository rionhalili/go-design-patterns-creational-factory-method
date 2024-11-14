package main

type strawberry struct {
	Cake
}

func newStrawberry() ICake {
	return &strawberry{
		Cake{
			name:     "Strawberry Cake",
			cakeType: "Strawberry",
		},
	}
}
