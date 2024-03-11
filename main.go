package main

import "fmt"

func main() {
	conferncename := "Go Conference"
	const confernceTicket = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T , remainingTickets is %T , ocnferenceName is %T\n", confernceTicket, remainingTickets, conferncename)
	fmt.Printf("Welcome to %v booking application\n", conferncename)
	fmt.Printf("We have totla of %v tickets and %v are still av", confernceTicket, remainingTickets)
	fmt.Println("Get your ticket  here to attend")

	var userName string
	var userTickets int
	//ask suer for thier name

	userName = "Tom"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
