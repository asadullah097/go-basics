package main

import "fmt"

func main () {
	var  conferenceName string = "Go Conference"
	const conferenceTickets int= 50
	var  remingTickets uint = 50
	
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName,)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remingTickets)
	fmt.Println("Get your tickets here!")

	var firstName string 
	var lastName string
	var email string
	var  userTickets uint
	var bookings []string
	//user inputs
	fmt.Printf("Enter your first name: \n")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: \n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: \n")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: \n")
	fmt.Scan(&userTickets)
    remingTickets =remingTickets - userTickets

	bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("Whole array %v\n", bookings)
	fmt.Printf("The length of array is %v\n", len(bookings))
	fmt.Printf("Type of array is %T\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remingTickets, conferenceName)
}