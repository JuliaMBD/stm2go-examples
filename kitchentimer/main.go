package main

import (
	"machine"
	"time"

	testpackage "github.com/example/testpackage/src"
)

type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(s ...interface{}) {
	// println(s)
}

func main() {
	b1 := NewButton(machine.BUTTON_3)
	b2 := NewButton(machine.BUTTON_2)
	b3 := NewButton(machine.BUTTON_1)
	disp := NewLCDDisplay(320, 240)
	buz := NewBuzzer()

	testpackage.ConfigureDevice(b1, b2, b3, disp, buz)
	testpackage.ConfigureLogger(logTest)

	tick := make(chan bool)

	go func() {
		for {
			tick <- true
			time.Sleep(1 * time.Millisecond)
		}
	}()

	for {
		select {
		case <-tick:
			testpackage.InputDevice()
			testpackage.EntryrootStm0Task()
			testpackage.OutputDevice()
		}
	}
}
