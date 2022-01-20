package main

import(
	"fmt"
	"time"
)

func main()  {
	players :=[]string{"Michael Jordan", "Scottie Pippen", "Dennis Rodman", "Steve Kerr", "Toni Kukoc"}
	firstRoaster :=make(chan string)

	//Bulucu
	go func(dizi[] string){
		for _, player := range dizi{
			firstRoaster <- player //kanala gonderiliyor
			time.Sleep(time.Second)
		}
	}(players)

	//Alici
	go func(){
		for i:=0; i<5; i++{
			bulunan:= <-firstRoaster //kanaldan aliyoruz
			fmt.Println("Alici: Bulucudan "+bulunan+" alindi")
		}
	}()
	<-time.After(time.Second*5)
}