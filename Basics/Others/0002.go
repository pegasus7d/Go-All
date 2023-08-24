package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Println("mystring is set to",myString)
	ChangeUsingPointer(&myString)
	log.Println("mystring is set to",myString)
}

func ChangeUsingPointer(s *string){
	newValue:="Red";
	*s=newValue
}