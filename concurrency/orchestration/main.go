package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	foo()
	bar()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	go foo()
	go bar()
	time.Sleep(time.Second)
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
}
