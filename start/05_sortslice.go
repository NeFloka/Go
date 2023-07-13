package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []int{1, 4, 3, 5, 2}
	sort.Ints(s)
	fmt.Println(s)
}
