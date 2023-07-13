package main

import "fmt"

func main() {

	a := 1986

	switch {

	case a >= 1946 && a < 1965:
		fmt.Println("Привет, бумер!")

	case a >= 1965 && a < 1981:
		fmt.Println("Привет, представитель X!")

	case a >= 1981 && a < 1997:
		fmt.Println("Привет, миллениал!")

	case a >= 1997 && a < 2013:
		fmt.Println("Привет, зумер!")

	case a >= 2013:
		fmt.Println("Привет, альфа!")

	default:
		fmt.Println("Здравствуйте!")
	}
}
