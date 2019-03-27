package hello

import "testing"

func TestHello(t *testing.T) {
	w := "Hello, world."
	if got := Hello(); got != w {
		t.Errorf("want %q, get %q", w, got)
	}
}
