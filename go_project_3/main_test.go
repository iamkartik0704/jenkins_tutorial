package main

import "testing"

func TestGreet(t *testing.T) {
	if got, want := greet(), "Hello Jenkins!"; got != want {
		t.Errorf("greet() = %q, want %q", got, want)
	}
}