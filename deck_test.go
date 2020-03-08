package main

import "testing"

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Got wrong length, expected 16, got %v", len(d))
	}
}
