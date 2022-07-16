package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16{
		t.Errorf("Expected deck 16, but got %v", len(d))
	}
	
	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card to be Ace of Spades, got %v", d[0])
	}
	
	if d[len(d)-1] != "Four of Clubs"{
		t.Errorf("Expected last card to be Four of Clubs, got %v", d[len(d)-1])
		}
}
