package booker

import "fmt"

func bookTickets(firstNames []string) {
	registerCount = registerCount - userTickets
	bookings[0] = firstName + " " + lastName + " " + email
	bookin = append(bookin, userData)

	fmt.Printf("%v Event has been fully booked\n", websiteName)

	fmt.Printf("Thank you for registering %v you have successfully booked %v tickets \nYou will receive an email @ %v", firstName, userTickets, email)
	fmt.Printf("\n%v tickets remaining for %v event\n", registerCount, websiteName)

	// fmt.Printf("\nContent of the slice %v\n", bookin)
	// fmt.Printf("%v\n", bookin[0])
	// fmt.Printf("the content of the array: %v\n", bookings)

	fmt.Printf("the slice type: %T\n", bookin)
	fmt.Printf("the number of bookers in slice : %v\n", len(bookin))

}
