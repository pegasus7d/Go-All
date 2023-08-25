package main

import (
	"errors"
	"fmt"
	"net/http"
)
const portnumber=":8080"
func Home(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"This is the Home Page")
}
func About(w http.ResponseWriter,r *http.Request){
	sum:=addValues(2,2)
	_, _=fmt.Fprintf(w,fmt.Sprintf("This is the about page and 2 + 2 is %d",sum))
}
func Divide(w http.ResponseWriter,r *http.Request){
	f,err:=dividevalues(100.0,0.0)
	if(err!=nil){
		fmt.Fprintf(w,"Cannot divide by zero")
		return
	}
	fmt.Fprintf(w,fmt.Sprintf("%f divided by %f:%f",100.0,10.0,f))
}
func dividevalues(x,y float32)(float32,error){
	if y==0{
		err:=errors.New("cannot divide by zero")
		return 0,err
	}
	result:=x/y 
	return result,nil
}
func addValues(x,y int)int{
	return x+y
}


func main() {
	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)
	http.HandleFunc("/divide",Divide)
	fmt.Println(fmt.Sprintf("Starting application on port %s",portnumber))
	_=http.ListenAndServe(portnumber,nil)

}