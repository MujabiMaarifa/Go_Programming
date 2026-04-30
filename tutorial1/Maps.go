package main
import (
	"fmt"
)

func main() {
	//maps - map[key type]valueType{key1:value1, key2:value2}
	var cars = map[string]string{"Brand": "Volvo", "Series": "EC40", "Model year": "2026", "Country of Origin": "Sweden"}
	fmt.Println(cars)

	//german machines
	var a = make(map[string]string) //this is an empty map
	a["Brand"] = "Volkswagen"
	a["Series"] = "ID.7"
	a["Model year"] = "2024"
	a["Country of Origin"]= "German"

	a1 := make(map[string]string)
	a1["Brand"] = "Mercedes Benz"
	a1["Series"] = "GLC"
	a1["Model year"] = "2016"
	a1["Country of Origin"] = "Germany"

	a2 := make(map[string]string)
	a2["Brand"] = "Audi"
	a2["Series"] = "Q6 e-tron"
	a2["Model year"] = "2024"
	a2["Country of origin"] = "Germany"

	var a3 = make(map[string]string)
	a3["Brand"] = "BMW"
	a3["Series"] = "iX"
	a3["Model year"] = "2022"
	a3["Country of origin"] = "Germany"

	var a4 = make(map[string]string)
	a4["Brand"] = "Porsche"
	a4["Series"] = "Taycan"
	a4["Model year"] = "2020"
	a4["Country of Origin"] = "German"

	//combining the german models using slices
	German := []map[string]string{a1, a2, a3, a4}
	for i, v := range German {
		fmt.Printf("Map %d: %v\n", i+1, v)
	}
}

// map - stores the key value pairs
// map - is simply an unordered key value pair store that is also changeable and allows duplicates
//default value is nil
//map holds references to hashtables


