package main

import "log"

func main() {
	n:=5
	if n%2==1 {
		log.Println("Odd")
	}else{
		log.Println("Even")
	}
	myVar:="cat"
	switch myVar{
	case "cat":
		log.Println("cat is cat")
	case "dog":
		log.Println("cat is dog")
	default:
		log.Println("cat is set to default")
	}

}