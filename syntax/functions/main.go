package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type person struct {
	first string
	last  string
	age   int
}

//func (r Reciever) funcName(arg argType) returns returntypes {}
func (p person) thisGuyReads(urls ...string) string {
	var pages string

	for _, url := range urls {
		res, _ := http.Get(url)
		page, _ := ioutil.ReadAll(res.Body)
		pages += string(page)
		res.Body.Close()
	}
	p.first = "li'l' Miss"
	return pages
}

var p1 person
var p2 *person

func init() {
	p1 = person{"James", "Bond", 20}
	p2 = &person{"John", "Wayne", 22}
}

func main() {
	read := p2.thisGuyReads("http://www.geekwiseacademy.com/")
	fmt.Printf("** %s **\n", p2.first)
	fmt.Printf("%s", read)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// Only slices, maps, and channels are already refereces
	read = p1.thisGuyReads("http://www.geekwiseacademy.com/")
	fmt.Printf("%s\n", read)
	fmt.Printf("-- %s", p1.first)
}
