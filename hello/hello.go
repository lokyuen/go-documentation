package main

import (
	"log"
	"fmt"
	"example.com/greetings"
)

// import "rsc.io/quote"

func main() {
	// fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}