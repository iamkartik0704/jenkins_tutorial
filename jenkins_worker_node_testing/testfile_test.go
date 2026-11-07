package main
import "testing"

func TestHello(t *testing.T){
	requirement:="hello from pi"
	received:=Hello()
	if requirement!=received{
		t.Errorf("Hello() = %q, want %q", &received,requirement)
	}
}