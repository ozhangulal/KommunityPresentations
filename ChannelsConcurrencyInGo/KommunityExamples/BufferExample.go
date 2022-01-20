package main

import (
	"fmt"
)

func main() {
	channel:=make(chan string, 4)
	channel<-"Baslangic"
	channel<-"Ara Sicak"
	channel<-"Ana Yemek"
	channel<-"Tatli"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}