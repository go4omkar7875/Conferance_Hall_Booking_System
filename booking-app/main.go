package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()
	//or bookings := []string{}

	//fmt.Printf("conferenceTickets is %T , remainingTickets is %T, conference is %T\n", conferenceTickets, remainingTickets, conferenceName)

	//%T is placeholder for a type variable

	//fmt.Printf("Welcome to %v booking application\n", conferenceName)

	//%V is variable reference

	//fmt.Printf("We have toatal of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	//fmt.Println("Get your tickets here to attend")

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first name of bookings are : %v\n", firstNames)

		//fmt.Printf("These are all our bookings : %v\n", bookings)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Try next time")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("First Name and Last Name you entered is too short")

		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")

		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is Invalid")
		}

	}
	wg.Wait()

}

func greetUser() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have toatal of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name

	//scan takes from user input - pass by reference

	fmt.Println("Enter the First name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter the Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email Id: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	var userData = UserData{

		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########## ")
	fmt.Printf("Sending ticket:\n %v \n to email address %v \n", ticket, email)
	fmt.Println("##########")

	wg.Done()
}
