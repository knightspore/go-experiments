package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"

	"github.com/MarkKremer/microphone"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
)

func main() {

	err := microphone.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer microphone.Terminate()

	for i := 0; i < 3; i++ {
		filename := RecordClip(i)
		go Whisper(filename)
	}

}

func Whisper(filename string) {
	_, err := exec.Command("whisper", filename, "--language", "en").Output()
	if err != nil {
		fmt.Printf("Whisper Err: %s", err)
	}
	PrintTranscription(filename)
}

func RecordClip(i int) string {

	file, err := CreateClipFile(i)
	if err != nil {
		log.Fatal(err)
	}

	stream, format, err := microphone.OpenDefaultStream(44100, 2)
	if err != nil {
		log.Fatal(err)
	}

	stream.Start()

	go func(w io.WriteSeeker, s beep.Streamer, format beep.Format) {
		err = wav.Encode(w, s, format)
		if err != nil {
			log.Fatal(err)
		}
	}(file, stream, format)

	time.Sleep(10 * time.Second)

	stream.Stop()
	err = stream.Close()
	if err != nil {
		fmt.Println("Error closing stream")
		panic(err)
	}

	return file.Name()

}
