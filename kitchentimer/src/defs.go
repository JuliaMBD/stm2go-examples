package testpackage

import (
	"fmt"
)

// Device interface
type Button interface {
	Get() bool
}

type Display interface {
	PrintClock(m, s uint8)
}

type Buzzer interface {
	Sing()
	Stop()
}

func ConfigureDevice(b1, b2, b3 Button, d1 Display, b Buzzer) {
	DeviceButton1 = b1
	DeviceButton2 = b2
	DeviceButton3 = b3
	DeviceDisplay = d1
	DeviceBuzzer = b
}

// global vars
var (
	count   uint16
	button1 bool
	button2 bool
	button3 bool
	min     uint8
	sec     uint8

	DeviceButton1 Button
	DeviceButton2 Button
	DeviceButton3 Button
	DeviceDisplay Display
	DeviceBuzzer  Buzzer
)

const (
	Sec1  = 1000
	MSec1 = 1
)

func init() {
	count = 0
	button1 = false
	button2 = false
	button3 = false
	min = 0
	sec = 0
}

func InputDevice() {
	button1 = DeviceButton1.Get()
	button2 = DeviceButton2.Get()
	button3 = DeviceButton3.Get()
}

func OutputDevice() {
	// DeviceDisplay.PrintClock(min, sec)
}

func decreaseSec() {
	if sec == 0 && min != 0 {
		sec = 59
		min -= 1
		if debug {
			logger.Println(fmt.Sprintf("Decreasing 1sec %d : %d", min, sec))
		}
	} else {
		sec -= 1
		if debug {
			logger.Println(fmt.Sprintf("Decreasing 1sec %d : %d", min, sec))
		}
	}
	DeviceDisplay.PrintClock(min, sec)
}

func increaseMin() {
	min = (min + 1) % 60
	if debug {
		logger.Println(fmt.Sprintf("Increasng 1min %d : %d", min, sec))
	}
	DeviceDisplay.PrintClock(min, sec)
}

func increaseSec() {
	sec = (sec + 1) % 60
	if debug {
		logger.Println(fmt.Sprintf("Increasng 1sec %d : %d", min, sec))
	}
	DeviceDisplay.PrintClock(min, sec)
}

func displayClock(m, s uint8) {
	DeviceDisplay.PrintClock(min, sec)
}

func sing() {
	DeviceBuzzer.Sing()
}

func stop() {
	DeviceBuzzer.Stop()
}
