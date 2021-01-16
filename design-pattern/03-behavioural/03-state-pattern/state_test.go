package state

import "testing"

func TestState(t *testing.T) {
	c := NewContext(NewStateA())
	if c.GetState() != "state A" {
		t.Error("test error")
	}
	c.Request()
	if c.GetState() != "state B" {
		t.Error("test error")
	}
	c.Request()
	if c.GetState() != "state A" {
		t.Error("test error")
	}
	c.Request()
	if c.GetState() != "state B" {
		t.Error("test error")
	}
}
