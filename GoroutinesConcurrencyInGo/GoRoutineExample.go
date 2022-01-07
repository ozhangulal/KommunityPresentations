package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	/*go*/ func() {
		for i := 4; i < 7; i++ {
			fmt.Println(i)
		}
	}()

	/*go*/
	func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Second)
}

//Yukarıdaki senaryoda, goanahtar kelimeyi kendi kendine çalışan işlevlere ekliyoruz . Yürütme mainişlevden başlar .
//
//goAnahtar kelimeyle karşılaştığı anda , uygulamaya başka bir Go iş parçacığı ekleyen ayrı bir Goroutine oluşturur.
//	Bu iş parçacığı, işlevi ayrı bir eşzamanlı evre üzerinde yürütmekten sorumludur.

//100ms den 7ms dustu
