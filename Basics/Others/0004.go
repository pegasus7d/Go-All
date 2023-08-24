package main

import "log"

type mystruct struct {
	Firstname string
}

func (m *mystruct) printFirstName() string{
	return m.Firstname
}

func main() {
	var myVar mystruct
	myVar.Firstname = "John"

	myVar2 := mystruct{
		Firstname: "Mary",
	}
	log.Println("myvar is set to",myVar.Firstname)
	log.Println("myvar2 is set to",myVar2.Firstname)
	log.Println("myvar is set to",myVar.printFirstName())
	log.Println("myvar2 is set to",myVar2.printFirstName())
}