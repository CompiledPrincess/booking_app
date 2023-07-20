package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, bookedTickets uint, remainingTickets uint) (bool, bool, bool) {
	isNameValid := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isBookedTicketsValid := bookedTickets > 0 && bookedTickets <= uint(remainingTickets)

	return isNameValid, isEmailValid, isBookedTicketsValid
}
