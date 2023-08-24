package main

import (
	"log"

	"github.com/pegasus7d/myprog/helpers"
)

const numPool=1000
func CalculateValue(intChan chan int) {
	RandomNumber := helpers.RandomNumber(numPool)
	intChan <- RandomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <- intChan

	log.Println(num)


}
