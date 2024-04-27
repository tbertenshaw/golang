package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message:= greetings.Hello("Timbo")
	fmt.Println(message)
}