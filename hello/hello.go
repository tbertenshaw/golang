package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Tim","Gladys","Fred"}

	messages,err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
