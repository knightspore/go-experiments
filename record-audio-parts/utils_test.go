package main

import "testing"

func TestFilename(t *testing.T) {
	want := "recording_0.wav"
	got := Filename(0)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
