package main

import "fmt"

func main() {
	her := avengers{"Batman", "Superman", "Spiderman"}
	her = append(her, "Ironman")
	her.assemble()
	fmt.Println(her)
}
