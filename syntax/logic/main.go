package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func switchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
		//fallthrough()  // fallthrough will invoke the next case statement
	case string:
		fmt.Println("string")
	case contact, func() contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")

	}
}

func main() {
	switchOnType(7)
	switchOnType("McLeod")
	var t = contact{"Good to see you,", "Tim"}
	var u = func() contact { return contact{"1", "2"} }
	switchOnType(t)
	switchOnType(u)

	if t.name == "John" {
		switchOnType(t.name)
	} else {
		switchOnType(t.greeting)
	}
}
