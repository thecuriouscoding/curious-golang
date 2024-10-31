package main

import (
	"fmt"
	"time"
)

var globalCounter int = 0

func increaseCounter() {
	fmt.Println("globalCounter before is: ", globalCounter)
	globalCounter++
	fmt.Println("globalCounter after is: ", globalCounter)
}

func SetInterval(t time.Duration, task func()) chan bool {
	stop := make(chan bool)
	ticker := time.NewTicker(t)
	go func() {
		for {
			select {
			case <-ticker.C:
				task()
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()
	return stop
}

func main() {
	stop := SetInterval(5*time.Second, increaseCounter)
	select {
	case <-time.After(20 * time.Second):
		stop <- true
	}
}
