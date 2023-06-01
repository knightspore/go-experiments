package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Filename(i int) string {
	return fmt.Sprintf("recording_%d.wav", i)
}

func PrintTranscription(filename string) {
	file, err := os.Open(fmt.Sprintf("%s.json", filename))
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

func CreateClipFile(i int) (*os.File, error) {
	filename := Filename(i)
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}
