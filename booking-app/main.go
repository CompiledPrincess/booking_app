package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Compiled Princess Conference"
	const conferenceTickets = 333
	var remainingTickets = 333
	var bookings = []string{}

	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v tickets are availible!\n", conferenceTickets, remainingTickets)
	fmt.Println("Grab your tickets here!")

	for {
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

		isNameValid := len(firstName) >= 2 && len(lastName) >= 2
		isEmailValid := strings.Contains(email, "@")
		isBookedTicketsValid := bookedTickets > 0 && bookedTickets <= uint(remainingTickets)

		if isNameValid && isEmailValid && isBookedTicketsValid {
			remainingTickets = remainingTickets - int(bookedTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v,you booked %v tickets. You will receive a confirmation email at %v \n",
				firstName, lastName, bookedTickets, email)

			fmt.Printf("The remaining tickets' total is %v \n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
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
