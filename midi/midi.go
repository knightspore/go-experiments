package main

import (
	"fmt"
	"time"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
)

func main() {
	defer midi.CloseDriver()

	in, err := midi.InPort(1)
	if err != nil {
		fmt.Println("can't find AKAI")
		return
	}

	stop, err := midi.ListenTo(in, func(msg midi.Message, timestampms int32) {
		var bt []byte
		var ch, key, vel uint8
		switch {
		case msg.GetSysEx(&bt):
			fmt.Printf("got sysex: % X\n", bt)
		case msg.GetNoteStart(&ch, &key, &vel):
			fmt.Printf("%s %s\n", midi.Note(key), Represent(vel))
		// case msg.GetNoteEnd(&ch, &key):
		// fmt.Printf("ending note %s on channel %v\n", midi.Note(key), ch)
		default:
			// ignore
		}
	}, midi.UseSysEx())

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	for {
		time.Sleep(time.Second * 5)
	}

	stop()
}

func Represent(v uint8) string {
	var s string

	for i := uint8(0); i < v; i++ {
		s += "."
	}

	return s
}
