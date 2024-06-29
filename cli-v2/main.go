package main

import (
	"flag"
	"fmt"
)

func main() {
	// go run main.go -name=salman
	firstName := flag.String("name", "example = Salman", "example = Taghooni")
	flag.Parse()
	fmt.Println("name is", *firstName)
}
