package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user
	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	//	outputUserDetails(appUser)
	outputUserDetails(&appUser) // using pointers

	//	appUser = user{ ## can define like this but ensure sequence should match with the definition above
	//		userFirstName,
	//		userLastName,
	//		userBirthdate,
	//		time.Now(),
	//	}
	// ... do something awesome with that gathered data!
}

//func outputUserDetails(u user) {
//	fmt.Println(u.firstName, u.lastName, u.birthDate)
//
//}

// using pointers
func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
