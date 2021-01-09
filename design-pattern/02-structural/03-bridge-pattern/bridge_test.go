package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {
	abs := NewAbstraction()
	abs.SetImplementor(&ConcreateImplementorA{})
	if abs.Operation() != "implementor a" {
		t.Error("TestBridge1 fail")
	}
	abs.SetImplementor(&ConcreateImplementorB{})
	if abs.Operation() != "implementor b" {
		t.Error("TestBridge2 fail")
	}
}
