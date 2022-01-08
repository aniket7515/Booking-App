package main

import "fmt"

func main() {
	var conferenceName = "Go conference"

	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName) // we can write the print function with the fmt
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get you tickets to attend")

}

// the error comes in go when we import package or declare variable but dont use it.
// we use const to the value that does not change
