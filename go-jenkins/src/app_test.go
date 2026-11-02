package main

import "testing"

func TestGreeting(t *testing.T) {
	want := "Hello from Go Application for Jenkins CI/CD!"
	got := Greeting()
	if got != want {
		t.Errorf("Greeting() = %q, want %q", got, want)
	}
}
