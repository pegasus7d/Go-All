package main

import (
	"fmt"
	"net/http"

	"github.com/pegasus7d/go-course/pkg/handlers"
)
const portnumber=":8080"




func main() {
	http.HandleFunc("/",handlers.Home)
	http.HandleFunc("/about",handlers.About)
	fmt.Printf(fmt.Sprintf("Starting application on port %s",portnumber))
	_=http.ListenAndServe(portnumber,nil)

}