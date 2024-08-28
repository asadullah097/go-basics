package main

import (
	"fmt"
	"strings"
)
	var conferenceName string = "Go Conference"
	const conferenceTickets int= 5
	var remingTickets uint = 5
    var bookings []string
func main () {
	
	greetUser()
	for {
		
	    firstName, lastName,email, userTickets := getUserInputs()
		//validations
		isValidName, isValidEmail, isValidTicketNumber :=validateUserInput(firstName, lastName, email, userTickets, remingTickets)
   		if isValidName && isValidEmail && isValidTicketNumber {
			bookings:=bookTicket(remingTickets, userTickets, firstName, lastName, email, bookings, conferenceName)
			var firstNames =getFirstNames(bookings)
			
			fmt.Printf("These all are in our bookings %v\n", firstNames)
			if remingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year\n")
				break
			}
	    }else {
			if !isValidName {
				fmt.Println("Invalid name, please try again")
			}
			if !isValidEmail {
				fmt.Println("Invalid email, please try again")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid ticket number, please try again")
			}
		}
    }
	
}
func greetUser()  {
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName,)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remingTickets)
	fmt.Println("Get your tickets here!")
}
func getUserInputs() (string, string, string, uint) {
	var firstName string 
	var lastName string
	var email string
	var  userTickets uint

	//user inputs
	fmt.Printf("Enter your first name: \n")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: \n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: \n")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: \n")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func  getFirstNames(bookings []string) []string {
	firstNames:=[]string{}
	for _,booking := range bookings {
		var names =strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(remingTickets uint,userTickets uint, firstName string, lastName string, email string, bookings []string, conferenceName string) []string {
	remingTickets =remingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)
			
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remingTickets, conferenceName)
	return bookings
}