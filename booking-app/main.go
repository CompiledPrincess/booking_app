package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 333

var conferenceName = "Compiled Princess Conference"
var remainingTickets uint = 333
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, bookedTickets := getUserInput()
		isNameValid, isEmailValid, isBookedTicketsValid := validateUserInput(firstName, lastName, email, bookedTickets)

		if isNameValid && isEmailValid && isBookedTicketsValid {
			bookTickets(bookedTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("The tickets are sold out!")
				break
			}
		} else {
			if !isNameValid {
				fmt.Println("Please enter an valid first or last name.")
			}
			if !isEmailValid {
				fmt.Println("The email you entered is invalid.")
			}
			if !isBookedTicketsValid {
				fmt.Println("The number of tickets you booked is invalid.")
			}

		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v tickets are availible!\n", conferenceTickets, remainingTickets)
	fmt.Println("Grab your tickets here!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func validateUserInput(firstName string, lastName string, email string, bookedTickets uint) (bool, bool, bool) {
	isNameValid := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isBookedTicketsValid := bookedTickets > 0 && bookedTickets <= uint(remainingTickets)

	return isNameValid, isEmailValid, isBookedTicketsValid
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, bookedTickets
}

func bookTickets(bookedTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - bookedTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v,you booked %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, bookedTickets, email)
	fmt.Printf("The remaining tickets' total is %v \n", remainingTickets)
}
