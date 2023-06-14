package main

import "fmt"

func main() {
	conferenceName := "Compiled Princess Conference"
	const conferenceTickets = 333
	var remainingTickets = 333

	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v tickets are availible!\n", conferenceTickets, remainingTickets)
	fmt.Println("Grab your tickets here!")

	var firstName string
	var lastName string
	var email string
	var bookedTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("How many tickets are you booking?")
	fmt.Scan(&bookedTickets)

	remainingTickets = remainingTickets - int(bookedTickets)

	fmt.Printf("Thank you %v %v,you booked %v tickets. You will receive a confirmation email at %v \n",
		firstName, lastName, bookedTickets, email)

	fmt.Printf("The remaining tickets' total is %v \n", remainingTickets)

}
