package main

import (
	"Travel_Booking/authentication"
	"Travel_Booking/router"
)

func main() {

	users, err := authentication.SetAuthenticationConfig()
	if err != nil {
		// Handle the error (e.g., log it or terminate the application)
		panic(err)
	}
	router.Router(users)

}
