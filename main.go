package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString("GPIO No should be specified.\n")
		os.Exit(1)
	}

	gpio := os.Args[1]

	fexport, err := os.Create("/sys/class/gpio/export")
	if err != nil {
		panic(err)
	}

	defer fexport.Close()

	fexport.WriteString(gpio)

	fdirection, err := os.Create(fmt.Sprintf("/sys/class/gpio/gpio%s/direction", gpio))

	defer fdirection.Close()

	fdirection.WriteString("out")

	fvalue, err := os.Create(fmt.Sprintf("/sys/class/gpio/gpio%s/value", gpio))
	if err != nil {
		panic(err)
	}

	defer fvalue.Close()

	if err := rpio.Open(); err != nil {
		panic(err)
	}

	defer rpio.Close()

	pin := rpio.Pin(7)
	pin.Input()

	fmt.Println("Ready")

	for {
		if pin.Read() == rpio.High {
			fvalue.WriteString("1")
			fmt.Println("Motion Detected!")
		} else {
			fvalue.WriteString("0")
		}

		time.Sleep(1 * time.Second)
	}
}
