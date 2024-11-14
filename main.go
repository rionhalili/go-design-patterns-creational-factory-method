package main

import "fmt"

func main() {
	appleCake, _ := getCake("Apple")
	orangeCake, _ := getCake("Orange")

	fmt.Println("Displayed cakes:")
	fmt.Println("====================================")
	report(appleCake)
	report(orangeCake)
	fmt.Println("====================================")
}

func report(cake ICake) {
	fmt.Printf(
		"Requested cake: %s is categorized as %s cake and costs %.2f",
		cake.getName(),
		cake.getCategory(),
		cake.getPrice(),
	)
	fmt.Println()
}
