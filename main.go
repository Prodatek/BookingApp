package main

import (
	"fmt"
	"strconv"
	"strings"
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
var bookin = make([]map[string]string, 0)

// var firstNames []string

var userData = make(map[string]string)

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
			userData["firstname"] = firstName
			userData["lastname"] = lastName
			userData["email"] = email
			userData["userticket"] = strconv.FormatUint(uint64(userTickets), 10)

			bookin = append(bookin, userData)
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


func getfirstName(bookin []map[string]string) []string {
	var firstNames []string
	for _, booking := range bookin {
		firstNames = append(firstNames, booking[firstName])
	// 	if firstName, exists := booking["firstname"]; exists {
    //         firstNames = append(firstNames, firstName)
	
	// }
}
	return firstNames
}
