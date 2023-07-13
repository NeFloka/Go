package main

import "fmt"

func main() {

	var weekTemp = [7]int{1, 4, 2, 5, 4, 6, 3}

	sumTemp := 0

	for _, temp := range weekTemp {
		sumTemp += temp
	}
	average := sumTemp / len(weekTemp)

	fmt.Println(average)
}
