package main

import (
	"fmt"
	"time"
)

func channel_timer_golang() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer2 is fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 is stopped")
	}
	time.Sleep(2 * time.Second)
	fmt.Println("All is sstoped")
}
