package main

type redVelvetCake struct {
	Cake
}

func newRedVelvetCake() ICake {
	return &redVelvetCake{
		Cake: Cake{
			name:     "Cherry Cake",
			cakeType: "Cherry",
		},
	}
}
