package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {
	facade := NewFacade()
	a := facade.MethodA()
	b := facade.MethodB()
	if a != "one two four" {
		t.Error("test facade error")
	}
	if b != "two four" {
		t.Error("test facade error")
	}
}
