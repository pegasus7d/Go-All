package main

import (
	"log"
	// "os/user"
	// "time"
)



type User struct{
	Firstname string 
	LastName string 
	// PhoneNumber string 
	// Age int 
	// BirthDate time.Time
}

func main() {

	user:= User{
		Firstname: "Debayan",
		LastName: "Biswas",
	}
	log.Println(user.Firstname,user.LastName)
	
}
