package main 

import "fmt"

func main(){
	fmt.Println("Hello World")
	var whatToSay string 
	var i int 
	whatToSay="Goodbye ,Cruel World"
	fmt.Println(whatToSay)
	i=5
	fmt.Println("i is set to:",i)
	whatwassaid,theotherthing:=saySomething()
	fmt.Println("The function returned",whatwassaid,theotherthing)
}

func saySomething()(string,string) {
	return "something","else"
}