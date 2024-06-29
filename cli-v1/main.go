package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//Get Args By os.Args
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Hello ", "No Name")
		log.Fatalf("%d", len(args))
	}
	name := args[1]
	fmt.Println("Hello ", name)
}
