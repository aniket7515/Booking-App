package main

import "fmt"

func main() {
	var conferenceName = "Go conference"

	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to ", conferenceName, " booking application") // we can write the print function with the fmt
	fmt.Println("We have total of ", conferenceTickets, " tickets and", remainingTickets, "are still available")
	fmt.Println("Get you tickets to attend")

}

// the error comes in go when we import package or declare variable but dont use it.
// we use const to the value that does not change
