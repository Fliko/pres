package main

import (
	"encoding/json"
	"fmt"
)

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

	// Marshalling
	type response1 struct {
		Page   int
		Fruits []string
	}

	type response2 struct {
		Page   int      `json:"page"`
		Fruits []string `json:"fruits"`
	}

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
}
