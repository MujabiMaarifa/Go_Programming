package main
import (
	"fmt"
)

type CarDiscovery struct{
	Make string
	Model string
	ModelCountry string
	ModelCode string
	ModelYear string
	Price string
}

func main() {
	var LR1, LR2 CarDiscovery
	LR1.Make = "Land Rover(4338)"
	LR1.Model = "Discovery 419"
	LR1.ModelCountry = "England"
	LR1.ModelCode = "HSE TD6"
	LR1.ModelYear = "2017"
	LR1.Price = "USD 32473"

	//Land rover 2
	LR2.Make = "Land Rover(439)"
	LR2.Model = "Discovery 2"
	LR2.ModelCountry = "London, England"
	LR2.ModelCode = "HSE SD4"
	LR2.ModelYear = "2017"
	LR2.Price = "USD 28208"

	printCarDiscovery(LR1)
	printCarDiscovery(LR2)
}

func printCarDiscovery(cars CarDiscovery) {
	fmt.Println("Make: ", cars.Make)
	fmt.Println("Model: ", cars.Model)
	fmt.Println("Country of Origin: ", cars.ModelCountry)
	fmt.Println("Model Code: ", cars.ModelCode)
	fmt.Println("Model year: ", cars.ModelYear)
	fmt.Println("Price: ", cars.Price)
}
