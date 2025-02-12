package helped

import "fmt"

func userInput(firstName string, lastName string, email string, userTickets uint) (string, string, string, uint) {
	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email")
	fmt.Scan(&email)
	fmt.Println("Please enter the number of tickets you want to book")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
