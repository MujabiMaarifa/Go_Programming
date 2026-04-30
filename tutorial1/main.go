package main
import "fmt"


func main() {
	fmt.Println("Hello world")	

	//variables
	var intNum int64 = 32767
	intNum = intNum + 10000
	fmt.Println(intNum)

	var floatNum float32= 123445667.9884237957
	fmt.Println(floatNum)

	var floatNum2 float32 = 10.1
	var intNum2 int32 = 134

	//perform operations among them -- type casting one variable to another
	var sum float32 = floatNum2 + float32(intNum2)
	fmt.Println("The sum is: ", sum)

	//strings
	var myString string = "hello"
	fmt.Println(myString)

	//length of a string calculated in utf
	fmt.Println(len(myString))

	var myRune rune = 'a'
	fmt.Println(myRune)

	myVar := "This is a short hand text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const pi float32 = 3.142
	fmt.Println("value of pi is declared a const as we will not change it: ", pi)
}
