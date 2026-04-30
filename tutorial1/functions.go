package main
import "fmt"

func main() {
	greeting := "Hello!"
	name := "Maarifa"
	var id string = "200848850"
	greetings(greeting, name, id)
}

func greetings(greet string, userName string, id string) {
	fmt.Println(greet + ", " + userName + ": user id: " + id)
}
