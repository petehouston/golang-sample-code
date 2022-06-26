package main

import "github.com/gen2brain/beeep"

func execBeep() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}
}

func execNotify() {
	err := beeep.Notify("Hello", "Nice to meet you", "info.png")
	if err != nil {
		panic(err)
	}
}

func execAlert() {
	err := beeep.Alert("Alert!", "Timeout! You need to take a break!", "warning.png")
	if err != nil {
		panic(err)
	}
}

func main() {

	execNotify()
}
