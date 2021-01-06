package simplefactory

import "testing"

func TestSimpleFactory(t *testing.T) {
	prodA := SimpleFactory(ProdTypA)
	if prodA.Describe() != "A" {
		t.Error("simple factory prod error.")
	}
	prodB := SimpleFactory(ProdTypB)
	if prodB.Describe() != "B" {
		t.Error("simple factory prod error.")
	}
}
