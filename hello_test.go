package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("hello() = %q,want %q", got, want)
	}
}

func TestCallWorld(t *testing.T) {
	t.Log(CallWorld())
}
