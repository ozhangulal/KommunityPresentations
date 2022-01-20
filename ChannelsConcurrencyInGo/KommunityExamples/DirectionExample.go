package main

import (
	"fmt"
)

func sender(channel chan<- string, message string)  {
	channel<-message
}

func comm(receiver <-chan string, sender chan<-string){
	message:=<-receiver
	sender<-message
}

func main()  {
	pippen :=make(chan string, 1)
	jordan :=make(chan string, 1)
	sender(pippen, "Pippen assits, Jordan scored!")
	comm(pippen, jordan)
	fmt.Println(<-jordan)
}
