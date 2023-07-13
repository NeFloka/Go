package main

import (
	"fmt"
	"reflect"
)

func main() {

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	s3 := []string{"1", "2", "3"}
	s4 := []int{1, 2, 3}

	fmt.Println(reflect.DeepEqual(s1, s2)) // false
	fmt.Println(reflect.DeepEqual(s1, s3)) // false
	fmt.Println(reflect.DeepEqual(s1, s4)) // true
}
