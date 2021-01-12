package flyweight

import (
	"testing"
)

func TestFlyweight(t *testing.T) {
	extrinsicstate := 22
	fac := NewFlyweightFactory()
	fx := fac.GetFlyweight("X")
	extrinsicstate--
	x := fx.Operation(extrinsicstate)
	if x != "concrete:21" {
		t.Error("test flyweight error")
	}

	fy := fac.GetFlyweight("Y")
	extrinsicstate--
	y := fy.Operation(extrinsicstate)
	if y != "concrete:20" {
		t.Error("test flyweight error")
	}

	uf := NewUnsharedConcreteFlyweight()
	extrinsicstate--
	u := uf.Operation(extrinsicstate)
	if u != "unshared:19" {
		t.Error("test flyweight error")
	}
}
