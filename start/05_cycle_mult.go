package main

import "fmt"

func main() {

	var height int

	fmt.Scan(&height)

	if height > 185 {
		fmt.Println("Высокий рост")

	} else if height >= 170 && height <= 185 {
		fmt.Println("Средний рост")

	} else {
		fmt.Println("Небольшой рост")
	}

	var x int

	fmt.Scan(&x)

	if x > 1000 {
		fmt.Println("Apple")

	} else if x >= 500 && x <= 1000 {
		fmt.Println("Samsung")

	} else {
		fmt.Println("Nokia с фонариком")
	}

	var x int
	var res int = 1

	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		res *= i
	}
	fmt.Println(res)
}
