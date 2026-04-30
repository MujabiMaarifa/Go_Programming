package main
import (
	"fmt"
	
)

func main() {
	var intArr [3] int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//method 2 initializing an array
	myArr := [3] int32{1, 2, 3}
	fmt.Println(myArr)

	//dot syntax
	dotArr := [...] int32{4, 5, 6}
	fmt.Println(dotArr)

	//slices - similar to arrays only that there sizes grow
	var intSlice []int32 = []int32{7,8,9}
	fmt.Println(intSlice)
	intSlice=append(intSlice, 7)
	fmt.Println(intSlice)

	//slice
	mySlice := []string{"Maarifa", "Mujabi", "Jitu", "David", "Babu"}
	fmt.Println(mySlice)
	myNewSlice := append(mySlice, "Ojiambo")
	fmt.Println(myNewSlice[5])

	//functions of the slice
	//cap() -> num of elements that the slice can grow or shrink
	//len() -> returns the number of elements in the slice
	fmt.Println(len(myNewSlice))
	fmt.Println(cap(myNewSlice))
	fmt.Println(myNewSlice)

	//creating a slice from an array
	units := [6]string{"AI and ML", "TOC", "CC", "IGS", "SE", "MOBS"}
	fmt.Println(units)
	unitSlice :=units[0:4]
	fmt.Println("slice created from the array units")

	//checking the capacity and the length of the slice
	fmt.Printf("units size=%d\n",len(units))
	fmt.Printf("units capacity=%d\n", cap(units))
	fmt.Println("Slices")
	fmt.Printf("slice size = %d\n", len(unitSlice))
	fmt.Printf("slice capacity =%d\n", cap(unitSlice))
	fmt.Println(unitSlice)

	fmt.Println("\nWe can therefore comfortably conclude that arrays have fixed size and slices can grow")

	//appending slices to other slice
	nameUnitsSlices := append(myNewSlice, unitSlice...)
	fmt.Println(nameUnitsSlices)
	fmt.Println(&nameUnitsSlices[6])
	fmt.Println(&nameUnitsSlices[7])
	fmt.Println(&nameUnitsSlices[5])

	//make(type, size, capacity) -> used to initiallize the sizes of the slices, maps and channels
	topics := make([]string, 3, 5)
	copy(topics, units[:len(units)-1])
	fmt.Println(topics)

	//required size of the slice and copying to the memory	
	fmt.Printf("Copied size= %d ", len(topics))
	fmt.Printf("Copied capacity= %d ", cap(topics))
}
