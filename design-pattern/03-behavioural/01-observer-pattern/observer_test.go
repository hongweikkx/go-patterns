package observer

import "testing"

func TestObserver(t *testing.T) {
	observerX := NewConcreteObserver("X")
	observerY := NewConcreteObserver("Y")
	observerZ := NewConcreteObserver("Z")
	sub := NewSubject()
	sub.Attach(observerX, observerY)
	sub.Notify()
	gX := observerX.Get()
	gY := observerY.Get()
	gZ := observerZ.Get()
	if !gX || !gY || gZ {
		t.Errorf("test Observer error, x:%v, y:%v, z:%v", gX, gY, gZ)
	}

}
