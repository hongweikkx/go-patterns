package factory

import "testing"

func TestFactoryMethod(t *testing.T) {
	prodA := (&AFactory{}).NewProd()
	if prodA.Describe() != "A" {
		t.Error("simple factory prod error.")
	}
	prodB := (&BFactory{}).NewProd()
	if prodB.Describe() != "B" {
		t.Error("simple factory prod error.")
	}
}
