package main

import (
	"log"

	"github.com/pegasus7d/myprog/helpers"
)

func main(){
	log.Println("Hello")
	var myvar helpers.SomeType
	myvar.TypeName="Debayan"
	myvar.TypeNumber=5
	log.Println(myvar.TypeName,myvar.TypeNumber)
	
	
}