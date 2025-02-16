package main

import (
	"fmt"
	// "strconv"
	"strings"
	"time"
	"sync"
)

// import "log"
// import "net/http"
// import "os"
// import "path/filepath"

var registerCount uint = 50

const websiteName = "GoNaNa"

var firstName string
var lastName string
var email string
var userTickets uint

// Array

var bookings [50]string

// Slice
var bookin = make([]UserData, 0)

// var firstNames []string

// var userData = make(map[string]string)

// using a struct in place of maps
type UserData struct{
	firstName string
	lastName string
	email string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := userInput()
		isvalidName, isvalidEmail, isvalidTickets := validateData(firstName, lastName, email, userTickets, registerCount)

		// only call array lists after they have been declared

		if isvalidEmail && isvalidName && isvalidTickets {

			registerCount = registerCount - userTickets
			bookings[0] = firstName + " " + lastName + " " + email

			//map of strings
			// userData["firstname"] = firstName
			// userData["lastname"] = lastName
			// userData["email"] = email
			// userData["userticket"] = strconv.FormatUint(uint64(userTickets), 10)

			// Userdata structure
			var userData = UserData {
				firstName: firstName,
				lastName: lastName,
				email: email,
				userTickets: userTickets,
			}

			bookin = append(bookin, userData)

			bookTicket(firstName, userTickets, email, registerCount,websiteName, bookin)
			

			// wg.Add(1)

			go sendTicket(userTickets, firstName, email)

			var noMoreTickets bool = registerCount == 0
			if noMoreTickets {

				fmt.Printf("We are fully books this year, dont miss out Next year ! ")
				break

			}
		} else {

			if !isvalidName {
				fmt.Printf("Name entered is too short\n")
			}
			if !isvalidEmail {
				fmt.Printf("email format invalid\n")
			}
			if !isvalidTickets {
				fmt.Printf("Invalid ticket amount\n")
			}

		}
}
		// wg.Wait()
}
func validateData(firstName string, lastName string, email string, userTickets uint, registerCount uint) (bool, bool, bool) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidTickets := userTickets > 0 && userTickets <= registerCount
	return isvalidName, isvalidEmail, isvalidTickets
}

func greetUsers() {
	fmt.Println("BOOKING APP 101")
	fmt.Printf("welcome to %v events Booking App %v Users tickets available\n", websiteName, registerCount)

	fmt.Printf("Welcome to %v booking application \nThe total number of bookings available is %v\n", websiteName, registerCount)
}


func getfirstName(bookin []UserData) []string {
	var firstNames []string
	for _, booking := range bookin {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}


func bookTicket(firstName string, userTIckets uint, email string, registerCount uint, websiteName string, bookin []UserData){
	fmt.Printf("the list of bookings are %v \n", bookin)

	// fmt.Printf("%v Event has been fully booked\n", websiteName)

	fmt.Printf("Thank you for registering %v you have successfully booked %v tickets \nYou will receive an email @ %v", firstName, userTickets, email)
	fmt.Printf("\n%v tickets remaining for %v event\n", registerCount, websiteName)

	// fmt.Printf("\nContent of the slice %v\n", bookin)
	// fmt.Printf("%v\n", bookin[0])
	// fmt.Printf("the content of the array: %v\n", bookings)

	fmt.Printf("The slice type: %T\n", bookin)
	fmt.Printf("the number of bookers in slice : %v\n", len(bookin))

	firstNames := getfirstName(bookin)

	fmt.Printf("the first names of bookers are: %v\n", firstNames)
}


func sendTicket(userTickets uint, firstName string, email string){
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("\n%v successfully booked for %v", userTickets, firstName)
	fmt.Println("*******************")
	fmt.Printf("Sending Ticket:%v, \nSent ticket to %v\n", ticket, email)
	fmt.Println("*******************")

	// wg.Done()


}

