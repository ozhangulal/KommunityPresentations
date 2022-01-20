package main

import (
	"fmt"
	"time"
)

func Fonksiyon1(channel chan string)  {
	fmt.Println("Fonksyion1 calistiriliyor")
	time.Sleep(time.Second*2)
	channel<-"Fonksiyon 1 tamamlandi"
}

func Fonksiyon2(channel chan string)  {
	fmt.Println("Fonksyion2 calistiriliyor")
	time.Sleep(time.Second*5)
	channel<-"Fonksiyon 2 tamamlandi"
}

func Fonksiyon3(channel chan string)  {
	fmt.Println("Fonksyion3 calistiriliyor")
	time.Sleep(time.Second*3)
	channel<-"Fonksiyon 3 tamamlandi"
}

func main(){
	channelOne:=make(chan string, 1)
	channelTwo:=make(chan string, 1)
	channelThree:=make(chan string, 1)

	go Fonksiyon1(channelOne)
	go Fonksiyon2(channelTwo)
	go Fonksiyon3(channelThree)

	for i:=0; i<3; i++{
		select {
		case messageOne:=<-channelOne:
			fmt.Println(messageOne)
			case messageTwo :=<-channelTwo:
			fmt.Println(messageTwo)
		case messageThree:=<-channelThree:
		fmt.Println(messageThree)

		}
	}
}