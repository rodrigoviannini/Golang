package main

import (
	"fmt"
)

func main() {
	// name, age := returnNameIdade()
	// fmt.Println("Meu nome é", name, "e tenho", age, "anos")

	_, age := returnNameIdade()
	fmt.Println("Meu nome é Rodrigo", "e tenho", age, "anos")
}

func returnNameIdade() (string, int) {
	name := "Rodrigo"
	age := 39
	return name, age
}
