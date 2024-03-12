package main

import "fmt"

func main() {
	conferncename := "Go Conference"
	const confernceTicket = 50
	var remainingTickets uint
	remainingTickets = 50

	fmt.Printf("conferenceTickets is %T , remainingTickets is %T , ocnferenceName is %T\n", confernceTicket, remainingTickets, conferncename)
	fmt.Printf("Welcome to %v booking application\n", conferncename)
	fmt.Printf("We have totla of %v tickets and %v are still av\n", confernceTicket, remainingTickets)
	fmt.Println("Get your ticket  here to attend")

	var bookings []string
	// bookings[0] = "Nana"
	// bookings[1] = "Zaid"

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask suer for thier name
	fmt.Println("Enter your first Name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last Name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email Address")
	fmt.Scan(&email)
	fmt.Println("Enter Number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The Whole Slice:%v\n", bookings)
	fmt.Printf("The First value:%v\n", bookings[0])
	fmt.Printf("Slice Type:%T\n", bookings)
	fmt.Printf("Slice length:%v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking the tickets %v , You will  receive a confiramtaions email at %v \n ", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferncename)
}
