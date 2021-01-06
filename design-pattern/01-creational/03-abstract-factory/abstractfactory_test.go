package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	prodA := (&AFactory{}).NewProd2()
	if prodA.Describe2() != "A2" {
		t.Error("simple factory prod error.")
	}
	prodB := (&BFactory{}).NewProd1()
	if prodB.Describe1() != "B1" {
		t.Error("simple factory prod error.")
	}
}
