package main

import "fmt"

func main() {

	var myArray [50]int
	fmt.Println(myArray)

	mySlice := []int{1, 3, 5, 7, 9, 11}
	idomaticSlice := make([]int, 5, 10)
	fmt.Printf("%T - %v\n", mySlice, mySlice)
	fmt.Println(idomaticSlice)

	myMap := map[string]int{
		"one": 0,
		"two": 1,
	}
	idomaticMap := make(map[string]int)
	idomaticMap["one"] = 1
	fmt.Println(myMap)
	fmt.Println(idomaticMap)
}
