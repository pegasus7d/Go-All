package main

import (
	"log"
	"sort"
)

type myUser struct{
	FirstName string 
	LastName string 
}

func main() {

	// myMap := make(map[string]string)
	// myMap["dog"] = "Samson"
	// myMap["other-dog"] = "Cassie"
	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])
	myMap:= make(map[string]myUser)
	me:= myUser{
		FirstName:"Deb",
		LastName:"Biswas",
	}
	myMap["me"]=me
	log.Println(myMap["me"].FirstName)



	var myString []string 
	var myint []int 
	myString=append(myString, "Debayan")
	myString=append(myString, "Deb")
	myString=append(myString, "Debu")
	myint=append(myint, 4)
	myint=append(myint, 1)
	myint=append(myint, 3)
	log.Println(myString)
	log.Println(myint)
	sort.Ints(myint)
	log.Println(myint)
	numbers:=[]int{1,2,3,4,5,6}
	log.Println(numbers)
	log.Println(numbers[0:2])
}