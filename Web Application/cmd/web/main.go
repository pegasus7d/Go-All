package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pegasus7d/go-course/pkg/config"
	"github.com/pegasus7d/go-course/pkg/handlers"
	"github.com/pegasus7d/go-course/pkg/render"
)
const portnumber=":8080"




func main() {
	var app config.AppConfig 
	tc,err:=render.CreateTemplateCache()
	if err!=nil{
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache=tc
	http.HandleFunc("/",handlers.Home)
	http.HandleFunc("/about",handlers.About)
	fmt.Printf(fmt.Sprintf("Starting application on port %s",portnumber))
	_=http.ListenAndServe(portnumber,nil)

}