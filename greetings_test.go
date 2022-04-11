package greetings

import (
	"testing"
)

func TestHello(t *testing.T) {
	if Hello("apple") != "Hi, apple. Welcome!" {
		t.Error("Hello error")
	}
}
