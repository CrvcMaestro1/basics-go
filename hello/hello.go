package main

import (
	"example.com/greetings"
	"fmt"
	"log"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello everyone")
	fmt.Println(quote.Go())
	fmt.Println("Sum of 6 and 12 is: ", sum(6, 12))
	fmt.Println("Multiplier of 5.5 and 1.5 is: ", multiplier(5.5, 1.5))
	// modules
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Hello function
	message, errMessage := greetings.Hello("Christian")
	if errMessage != nil {
		log.Fatal(errMessage)
	}
	fmt.Println(message)
	// Hellos function
	names := []string{"Christian", "Andrea", "Anahi"}
	messages, errMessages := greetings.Hellos(names)
	if errMessages != nil {
		log.Fatal(errMessages)
	}
	fmt.Println(messages)
}

func sum(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func multiplier(firstNumber float32, secondNumber float32) float32 {
	return firstNumber * secondNumber
}
