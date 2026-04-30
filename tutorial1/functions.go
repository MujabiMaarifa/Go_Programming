package main
import (
	"fmt"
	"errors"
)

func main() {
	greeting := "Hello!"
	name := "Maarifa"
	var id string = "200848850"
	greetings(greeting, name, id)

	//performing operations with go
	fmt.Println("Arithmetic operations")

	var num int = 11
	var den int = 0 
	var quotient, remainder, err= intDivision(num, den)
	if err!=nil{
		fmt.Printf(err.Error())
	}else if remainder == 0{
		fmt.Printf("The result of the integer division is %v", quotient)
	}else{
		fmt.Printf("The qiotient of the integer division is %v, and the remainder is: %v ", quotient, remainder)
	}
}

func greetings(greet string, userName string, id string) {
	fmt.Println(greet + ", " + userName + ": user id: " + id)
}

func intDivision(numerator int, denomenator int) (int, int, error) {
	var err error
	//check if the denominator is 0
	if denomenator ==0 {
		err = errors.New("Cannot Divide by zero")
		return 0, 0, err
	}
	var quotient int = numerator/denomenator
	var remainder int = numerator % denomenator

	return quotient, remainder, err
}
