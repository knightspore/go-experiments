package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/MarkKremer/microphone"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
)

type Transcription struct {
	Text     string `json:"text"`
	Segments []struct {
		ID               int     `json:"id"`
		Seek             int     `json:"seek"`
		Start            float64 `json:"start"`
		End              float64 `json:"end"`
		Text             string  `json:"text"`
		Tokens           []int   `json:"tokens"`
		Temperature      float64 `json:"temperature"`
		AvgLogprob       float64 `json:"avg_logprob"`
		CompressionRatio float64 `json:"compression_ratio"`
		NoSpeechProb     float64 `json:"no_speech_prob"`
	} `json:"segments"`
	Language string `json:"language"`
}

func main() {

	for i := 0; i < 3; i++ {
		RecordClip(i)
		Whisper(filename(i))
	}

}

func Whisper(filename string) {
	fmt.Println("whisper")
	_, err := exec.Command("whisper", filename, "--language en").Output()
	if err != nil {
		fmt.Printf("Whisper Err: %s", err)
	}
	printTranscription(filename)
}

func printTranscription(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	var t Transcription
	json.Unmarshal(bytes, &t)
	fmt.Println(t.Text)
}

func filename(i int) string {
	return fmt.Sprintf("recording_%d.wav", i)
}

func createClipFile(i int) (*os.File, error) {
	filename := filename(i)
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func RecordClip(i int) {

	file, err := createClipFile(i)
	if err != nil {
		log.Fatal(err)
	}
	err = microphone.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer microphone.Terminate()

	stream, format, err := microphone.OpenDefaultStream(44100, 2)
	if err != nil {
		log.Fatal(err)
	}

	stream.Start()
	fmt.Println("record start")

	go func(w io.WriteSeeker, s beep.Streamer, format beep.Format) {
		err = wav.Encode(w, s, format)
		if err != nil {
			log.Fatal(err)
		}
	}(file, stream, format)

	time.Sleep(5 * time.Second)

	stream.Stop()
	fmt.Println("record stop")
	err = stream.Close()
	if err != nil {
		panic(err)
	}

}
