package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone,omitempty"`
}

// type Person struct {
// 	Name  string
// 	Age   int
// 	Phone string
// }

func main() {
	p := Person{
		Name: "khanhdev",
		Age:  23,
	}
	// p := Person{
	// 	Name:  "khanhdev",
	// 	Age:   23,
	// 	Phone: "0336849233",
	// }
	b, e := json.Marshal(p)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(string(b))
}
