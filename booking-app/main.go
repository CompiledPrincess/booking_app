package main

import "fmt"

func main() {
	var conferenceName = "Compiled Princess Conference"
	const conferenceTickets = 333
	var remainingTickets = 333

	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v tickets are availible!\n", conferenceTickets, remainingTickets)
	fmt.Println("Grab your tickets here!")

}
