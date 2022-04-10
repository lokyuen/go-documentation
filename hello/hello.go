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
	// message, err := greetings.Hello("Gladys")

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}

