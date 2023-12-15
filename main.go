package main

import (
	"fmt"
	"strings"
	"time"
)

const participantTickets int = 100

var eventCentre = "Etihad Stadium"
var remainingTickets uint = 100
var bookings = make([]Participant, 0)

type Participant struct {
	firstName       string
	middleName      string
	lastName        string
	email           string
	numberOfTickets uint
}

type Academy struct {
	firstName   string
	middleName  string
	lastName    string
	email       string
	age         int
	gender      string
	phoneNumber int
	tryOutDate  uint
}

// This is an event management system for football stadium's matches.
func main() {

	greetParticipant()

	var choice string
	fmt.Scan(&choice)

	switch choice {
	case "1":
		for {

			firstName, middleName, lastName, email, participantTickets := getParticipantInput()
			isValidName, isValidEmail, isValidTicketNumber := validateParticipantInput(firstName, middleName, lastName, email, participantTickets)

			if isValidName && isValidEmail && isValidTicketNumber {

				bookTicket(participantTickets, firstName, middleName, lastName, email)
				sendTicket(participantTickets, firstName, middleName, lastName, email)

				if remainingTickets == 0 {
					//end program
					fmt.Printf("Our %v booked out. Come back next week.", eventCentre)
					break
				}
			} else {
				if !isValidName {
					fmt.Println("All names shoul be at least 3 characters long.")
				}
				if !isValidEmail {
					fmt.Println("Email adrress you entered is not @gmail.com")
				}
				if !isValidTicketNumber {
					fmt.Println("Number of tickets is invalid")
				}
			}
		}
	case "2":
		for {

			firstName, middleName, lastName, email, age, gender, phoneNumber := getAcademyInput()
			isValidName, isValidEmail,isValideAge, isValidTryOutDate := validateParticipantInput(firstName, middleName, lastName, email, age, tryOutDate)

			if isValidName && isValidEmail && isValidTryOutDate {

				bookTicket(tryOutDate, firstName, middleName, lastName, email)
				sendTicket(tryOutDate, firstName, middleName, lastName, email)

				if remainingTickets == 0 {//try out date are full then end program
					fmt.Printf("Our %v booked out. Come back next week.", eventCentre)
					break
				}
			} else {
				if !isValidName {
					fmt.Println("All names shoul be at least 3 characters long.")
				}
				if !isValidEmail {
					fmt.Println("Email adrress you entered is not @gmail.com")
				}
				if !isValidAge >= 32 && <=15 {
					fmt.Println("Age you entered is below our benchmark")
				}
				if !isValidTicketNumber {
					fmt.Println("Number of tickets is invalid")
				}
			}
		}

	case "3":

	default:
		fmt.Println("Invalid choice")
	}

}

func greetParticipant() {
	fmt.Printf("Welcome to Etihad Stadium. Where we watch young stars play football to their hearts content.")
	fmt.Printf("To watch our young stars, Please press 1 to book a ticket. To register with an academy, press 2. If you're not a participant, press 3.")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func getParticipantInput() (string, string, string, string, uint) {
	var firstName string
	var middleName string
	var lastName string
	var email string
	var participantTickets uint

	//ask user for participant name

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your other name: ")
	fmt.Scan(&middleName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your userTickets: ")
	fmt.Scan(&participantTickets)

	return firstName, middleName, lastName, email, participantTickets
}

func validateParticipantInput(firstName string, middleName string, lastName string, email string, participantTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@gmail.com")
	isValidTicketNumber := participantTickets > 0 && participantTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(participantTickets uint, firstName string, middleName string, lastName string, email string) {
	remainingTickets = remainingTickets - participantTickets

	//create a map for a user
	var Participant = Participant{
		firstName:       firstName,
		middleName:      middleName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: participantTickets,
	}

	bookings = append(bookings, Participant)
	fmt.Printf("List of bookings is %v\n", bookings)

	//send confirmation of booking
	fmt.Printf("Dear %v %v %v, You have successfully booked %v tickets for the %v. Please check %v for confirmation\n", firstName, middleName, lastName, participantTickets, eventCentre, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, eventCentre)
}

func sendTicket(participantTickets uint, firstName string, middleName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v %v", participantTickets, firstName, middleName, lastName)
	fmt.Println("########")
	fmt.Printf("Sending ticket %v \nto email address %v\n", ticket, email)
	fmt.Println("########")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func getAcademyInput() (string, string, string, string, uint) {
	var firstName string
	var middleName string
	var lastName string
	var email string
	var age int
	var gender string
	var phoneNumber int
	var tryOutDate string

	//ask user for participant name

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your other name: ")
	fmt.Scan(&middleName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Enter your gender: ")
	fmt.Scan(&gender)

	fmt.Println("Enter your mobile number: ")
	fmt.Scan(&phoneNumber)

	return firstName, middleName, lastName, email, age, gender, phoneNumber
}

func validateAcademyInput(firstName string, middleName string, lastName string, email string, age int, tryOutDate string) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@gmail.com")
	isValidTicketNumber := participantTickets > 0 && participantTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func outDate(){

	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	fmt.Println("The time is", theTime)

	fmt.Println(theTime.Format("2006-1-2 15:4:5"))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// list of participants who booked for events
func listOfEvents() []string {
	names := []string{}
	for _, booking := range bookings {
		names = append(names, booking.names)
	}
	return names
}
