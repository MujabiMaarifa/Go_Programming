package main
import (
	"fmt"
)
func main(){
	//for initializer; condition; increment{} ->this loop creates and prints 5 numbers using 
	for i:=0; i<5; i++{
		//skips the value of 3
		if i==3 {
			continue
		}
		fmt.Println(i)
	}
	fruits()
	loopArray()
}

func fruits() {
	size := [2]string{"big", "tasty"}
	matunda := [3]string{"apple", "mango", "ndizi"}

	for i:=0; i<len(size); i++{
		for j:=0; j<len(matunda); j++{
			fmt.Println(size[i], matunda[j])
		}
	}

	fmt.Println("Assuming the index now we loop through using _ for the index place")
	for _, m :=range matunda{
		fmt.Printf("Matunda name: %v\n", m)
	}
}

func loopArray() {
	var nums = [5]int32{1, 2, 3, 4, 5}
	//using range returns index and the value of an array
	for i, v := range nums{
		fmt.Printf("Index: %v\t has stored value: %v\n", i, v) 
	}
}
