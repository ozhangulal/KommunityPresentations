
package main

import (
	"fmt"
	"sync"
	"time" 
)


func fonksiyon1(wg *sync.WaitGroup) {

	time.Sleep(2 * time.Second)
	fmt.Println("Fonk1 tamamlandı")


	wg.Done()
}

func fonksiyon2(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("Fonk2 tamamlandı")

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go fonksiyon1(&wg)
	go fonksiyon2(&wg)
	fmt.Println("Merhaba Dünya!")

	wg.Wait()

	fmt.Println("WaitGroup'lar tamamlandı.")
}