package mediator

import "testing"

func TestMediator(t *testing.T) {
	mediator := NewConcreteMediator()
	a := NewColleague("A", mediator)
	b := NewColleague("B", mediator)
	mediator.Colleague1 = a
	mediator.Colleague2 = b

	if a.Send("hello") != "B:hello" {
		t.Error("test error")
	}
}
