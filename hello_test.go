package main

import "testing"

func TestFoo(t *testing.T) {
	resultString := foo("Dude")
	if resultString != `"Shiiiii" KidneyDude` {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", resultString, `"Shiiiii" KidneyDude`)
	}
}
