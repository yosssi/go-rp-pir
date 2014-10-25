package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	defer rpio.Close()

	pin := rpio.Pin(7)
	pin.Input()

	fmt.Println("Ready")

	for {
		if pin.Read() == rpio.High {
			fmt.Println("Motion Detected!")
		}

		time.Sleep(1 * time.Second)
	}
}
