package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func timer() {
	timerCh := time.Tick(time.Duration(1) * time.Second)
	for range timerCh {
		fmt.Println("Sono le ore", time.Now())
	}
}

func main() {

	quit := make(chan struct{}, 0)

	stopButton := js.Global().Get("document").Call("getElementById", "stop")
	stopButton.Set("disabled", false)
	stopButton.Set("onclick", js.NewCallback(func([]js.Value) {
		println("stop al tempo!")
		stopButton.Set("disabled", true)
		quit <- struct{}{}
	}))

	go timer()

	<-quit
}
