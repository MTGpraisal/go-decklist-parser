package godecklistparser

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

// Card represents a parsed decklist card
type Card struct {
	Num             int
	Name            string
	Set             string
	CollectorNumber int
}
