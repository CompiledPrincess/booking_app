package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, bookedTickets uint) (bool, bool, bool) {
	isNameValid := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isBookedTicketsValid := bookedTickets > 0 && bookedTickets <= uint(remainingTickets)

	return isNameValid, isEmailValid, isBookedTicketsValid
}
