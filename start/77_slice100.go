package main

import "fmt"

func main() {

	hun := 100
	s := make([]int, 0, hun)

	for i := 0; i < hun; i++ {
		s = append(s, i+1)
	}
	fmt.Println(s)
}
